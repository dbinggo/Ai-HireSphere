package resume

import (
	"net/http"

	"Ai-HireSphere/application/interview/interfaces/api/internal/logic/resume"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetResumeFolderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := resume.NewGetResumeFolderListLogic(r.Context(), svcCtx)
		resp, err := l.GetResumeFolderList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
