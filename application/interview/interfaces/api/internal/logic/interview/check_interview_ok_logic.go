package interview

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"
	"Ai-HireSphere/application/interview/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckInterviewOkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCheckInterviewOkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckInterviewOkLogic {
	return &CheckInterviewOkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckInterviewOkLogic) CheckInterviewOk(req *types.CheckInterviewReq) (resp *types.CheckInterviewResp, err error) {
	// todo: add your logic here and delete this line

	return
}
