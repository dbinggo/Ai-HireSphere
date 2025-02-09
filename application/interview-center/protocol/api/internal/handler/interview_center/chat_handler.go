package interview_center

import (
	"Ai-HireSphere/application/interview-center/protocol/api/internal/logic/interview_center"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/svc"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/types"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ChatRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := interview_center.NewChatLogic(r.Context(), svcCtx)
		chat, err := l.Chat(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			w.Header().Set("Content-Type", "text/event-stream")
			w.Header().Set("Cache-Control", "no-cache")
			w.Header().Set("Connection", "keep-alive")
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.WriteHeader(200)
			controller := http.NewResponseController(w)
			for {
				select {
				case msg := <-chat:
					l.Debugf("Chat message: %v", msg)
					//if strings.Contains(msg, "[DONE]") {
					//	httpx.OkJsonCtx(r.Context(), w, struct {
					//		Data string `json:"data"`
					//	}{Data: msg})
					//	return
					//}

					//_, err := w.Write([]byte(msg))
					msg += "\n"
					_, err := fmt.Fprint(w, msg)
					if err != nil {
						fmt.Printf("Error writing response: %v", err)
						httpx.ErrorCtx(r.Context(), w, err)
						return
					}
					err = controller.Flush()
					if err != nil {
						fmt.Printf("Error flushing response: %v", err)
						httpx.ErrorCtx(r.Context(), w, err)
						return
					}
				}
			}
		}
	}
}
