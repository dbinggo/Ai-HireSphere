package interview

import (
	"net/http"

	"Ai-HireSphere/application/interview/interfaces/api/internal/logic/interview"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatNewHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := interview.NewChatNewLogic(r.Context(), svcCtx)
		resp, err := l.ChatNew()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
