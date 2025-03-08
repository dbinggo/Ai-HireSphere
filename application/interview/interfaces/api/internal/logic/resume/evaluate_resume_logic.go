package resume

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EvaluateResumeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEvaluateResumeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EvaluateResumeLogic {
	return &EvaluateResumeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EvaluateResumeLogic) EvaluateResume(req *types.EvaluateResumeReq) error {
	// todo: add your logic here and delete this line

	return nil
}
