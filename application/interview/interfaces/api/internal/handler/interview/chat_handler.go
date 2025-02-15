package interview

import (
	"net/http"

	"Ai-HireSphere/application/interview/interfaces/api/internal/logic/interview"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ChatHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SSEReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := interview.NewChatLogic(r.Context(), svcCtx)
		err := l.Chat(&req, w)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		}
	}
}
