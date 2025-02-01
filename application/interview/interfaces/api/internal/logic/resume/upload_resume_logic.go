package resume

import (
	"context"
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
func (l *UploadResumeLogic) UploadResume(r *http.Request) (resp *types.UploadResumeResp, err error) {
	// 拿到简历文件
	err = r.ParseMultipartForm(MAX_RESUME_SIZE)
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	file, header, err := r.FormFile("resume")
	if err != nil {
		l.Logger.Error(err)
		return nil, err
	}
	defer file.Close()
	url, err := l.svcCtx.ResumeAPP.UploadResume(l.ctx, file, header)

	return &types.UploadResumeResp{
		Address: url,
	}, err
}
