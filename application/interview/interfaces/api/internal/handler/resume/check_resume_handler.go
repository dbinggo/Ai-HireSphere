package resume

import (
	"Ai-HireSphere/common/ssex"
	"net/http"
	"strconv"
	"strings"
	"time"

	"Ai-HireSphere/application/interview/interfaces/api/internal/logic/resume"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CheckResumeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckResumeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		ctx := r.Context()
		l := resume.NewCheckResumeLogic(ctx, svcCtx)
		stream, err := l.CheckResume(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		sse := ssex.Upgrade(ctx, w)
		defer sse.Close()
		for {
			timer := time.NewTimer(time.Minute)
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-stream:
				var event ssex.SseEvent
				event.Event = msg.Event
				event.Data = msg.Data
				event.Id = strconv.Itoa(msg.ID)
				if !strings.Contains(event.Event, "Done") {
					sse.WriteEvent(event)
				}
				if !ok {
					return
				}
			case <-timer.C:
				return
			}
		}
	}
}
