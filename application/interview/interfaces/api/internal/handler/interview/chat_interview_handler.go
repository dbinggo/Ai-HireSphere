package interview

import (
	"Ai-HireSphere/common/ssex"
	"net/http"
	"time"

	"Ai-HireSphere/application/interview/interfaces/api/internal/logic/interview"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatInterviewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatInterviewReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		ctx := r.Context()
		l := interview.NewChatInterviewLogic(ctx, svcCtx)
		stream, err := l.ChatInterview(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		sse := ssex.Upgrade(ctx, w)
		defer sse.Close()
		for {
			timer := time.NewTimer(time.Minute) //每次流式输出只等待1分钟
			select {
			case <-ctx.Done():
				return
			case msg, ok := <-stream:
				var event ssex.SseEvent
				event.Event = msg.Event
				event.Data = msg.Data
				sse.WriteEvent(event)
				if !ok {
					return
				}
			case <-timer.C:
				return
			}
		}
	}
}
