package logic

import (
	"context"

	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/svc"
	"Ai-HireSphere/common/call/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserByIdLogic {
	return &FindUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserByIdLogic) FindUserById(in *types.Id) (*types.UserInfo, error) {
	// todo: add your logic here and delete this line

	return &types.UserInfo{}, nil
}
