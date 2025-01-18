package handler

import (
	"net/http"

	"Ai-HireSphere/application/user-center/interfaces/api/internal/logic"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CaptchaSendHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaSendReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewCaptchaSendLogic(r.Context(), svcCtx)
		err := l.CaptchaSend(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
