package interview_center

import (
	"net/http"

	"Ai-HireSphere/application/interview-center/protocol/api/internal/logic/interview_center"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/svc"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ResumeAddHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ResumeAddRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := interview_center.NewResumeAddLogic(r.Context(), svcCtx)
		resp, err := l.ResumeAdd(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
