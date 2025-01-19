package base

import (
	"Ai-HireSphere/common/model/enums"
	"context"

	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CaptchaVerifyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCaptchaVerifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CaptchaVerifyLogic {
	return &CaptchaVerifyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CaptchaVerifyLogic) CaptchaVerify(req *types.CaptchaVerifyReq) error {
	return l.svcCtx.BaseApp.CaptchaCheck(l.ctx, enums.CaptchaWayType(req.Way), req.Target, req.Code)
}
