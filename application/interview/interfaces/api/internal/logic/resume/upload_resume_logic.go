package resume

import (
	"Ai-HireSphere/common/codex"
	"context"
	"github.com/dbinggo/gerr"
	"net/http"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadResumeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadResumeLogic {
	return &UploadResumeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const MAX_RESUME_SIZE = 1024 * 1024 * 20 // 20M

// 简历上传服务
func (l *UploadResumeLogic) UploadResume(r *http.Request, req *types.UploadReusmeReq) (resp *types.UploadResumeResp, err error) {
	// 拿到简历文件
	err = r.ParseMultipartForm(MAX_RESUME_SIZE)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	if req.FolderId == 0 {
		return nil, gerr.Wraps(codex.FolderNameIsEmpty)
	}

	file, header, err := r.FormFile("resume")
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	defer file.Close()
	url, err := l.svcCtx.ResumeAPP.UploadResume(l.ctx, file, header, req.FolderId)

	return &types.UploadResumeResp{
		Address: url,
	}, err
}
