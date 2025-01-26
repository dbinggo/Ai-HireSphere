package interview_center

import (
	"context"

	"Ai-HireSphere/application/interview-center/protocol/api/internal/svc"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResumeAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewResumeAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResumeAddLogic {
	return &ResumeAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResumeAddLogic) ResumeAdd(req *types.ResumeAddRequest) (resp *types.ResumeAddResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
