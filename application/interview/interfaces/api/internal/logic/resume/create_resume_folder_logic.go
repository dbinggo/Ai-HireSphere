package resume

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateResumeFolderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateResumeFolderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateResumeFolderLogic {
	return &CreateResumeFolderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateResumeFolderLogic) CreateResumeFolder(req *types.CreqteResumeFolderReq) error {
	// todo: add your logic here and delete this line

	return nil
}
