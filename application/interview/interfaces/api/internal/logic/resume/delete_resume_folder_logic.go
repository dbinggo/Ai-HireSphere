package resume

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteResumeFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteResumeFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteResumeFolderLogic {
	return &DeleteResumeFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteResumeFolderLogic) DeleteResumeFolder(req *types.DeleteResumeFolderReq) error {
	return l.svcCtx.ResumeAPP.DeleteFolder(l.ctx, req.FolderId)
}
