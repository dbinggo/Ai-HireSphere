package resume

import (
	"Ai-HireSphere/common/ssex"
	"net/http"
	"strconv"
	"time"

	"Ai-HireSphere/application/interview/interfaces/api/internal/logic/resume"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func EvaluateResumeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EvaluateResumeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		ctx := r.Context()
		l := resume.NewEvaluateResumeLogic(ctx, svcCtx)
		stream, err := l.EvaluateResume(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		sse := ssex.Upgrade(ctx, w)
		defer sse.Close()
		for {
			timer := time.NewTimer(time.Minute * 5)
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-stream:
				var event ssex.SseEvent
				event.Event = msg.Event
				event.Data = msg.Data
				event.Id = strconv.Itoa(msg.ID)
				sse.WriteEvent(event)
				if !ok || msg.Event == "done" {
					return
				}
			case <-timer.C:
				return
			}
		}
	}
}
