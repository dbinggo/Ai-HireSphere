package resume

import (
	"net/http"

	"Ai-HireSphere/application/interview/interfaces/api/internal/logic/resume"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UploadResumeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadResumeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := resume.NewUploadResumeLogic(r.Context(), svcCtx)
		resp, err := l.UploadResume(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
