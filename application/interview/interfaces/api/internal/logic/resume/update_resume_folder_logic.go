package resume

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateResumeFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateResumeFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateResumeFolderLogic {
	return &UpdateResumeFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateResumeFolderLogic) UpdateResumeFolder(req *types.UpdateResumeFolderReq) error {

	return l.svcCtx.ResumeAPP.UpdateFolder(l.ctx, req.FolderId, req.FolderName)
}
