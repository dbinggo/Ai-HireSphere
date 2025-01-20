package logic

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InterviewLogic {
	return &InterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InterviewLogic) Interview(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
