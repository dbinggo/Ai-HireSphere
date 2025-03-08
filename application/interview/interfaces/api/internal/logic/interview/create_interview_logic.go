package interview

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateInterviewLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateInterviewLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateInterviewLogic {
	return &CreateInterviewLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateInterviewLogic) CreateInterview(req *types.NewInterviewReq) (resp *types.NewInterviewResp, err error) {
	// todo: add your logic here and delete this line

	return
}
