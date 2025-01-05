package logic

import (
	"Ai-HireSphere/application/user-center/interfaces/api/internal/dto"
	"context"

	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line

	way, user := dto.RegisterReq(req)

	res, err := l.svcCtx.UserApp.RegisterUser(l.ctx, way, user)

	if err != nil {
		return nil, err
	}

	return &types.RegisterResp{
		Token:  res,
		UserId: user.Id,
	}, nil

}
