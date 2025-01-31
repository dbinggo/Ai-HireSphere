package resume

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteResumeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteResumeLogic {
	return &DeleteResumeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteResumeLogic) DeleteResume(req *types.DeleteResumeReq) error {
	err := l.svcCtx.ResumeAPP.DeleteResume(l.ctx, req.ResumeId)
	return err
}
