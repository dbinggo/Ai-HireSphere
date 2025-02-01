package base

import (
	"net/http"

	"Ai-HireSphere/application/user-center/interfaces/api/internal/logic/base"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CaptchaVerifyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CaptchaVerifyReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := base.NewCaptchaVerifyLogic(r.Context(), svcCtx)
		err := l.CaptchaVerify(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
