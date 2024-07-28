package handler

import (
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"

	"Ai-HireSphere/app/api/internal/logic"
	"Ai-HireSphere/app/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		logx.Info("ping")

		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
