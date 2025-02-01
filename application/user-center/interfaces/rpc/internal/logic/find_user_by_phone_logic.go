package logic

import (
	"context"

	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/svc"
	"Ai-HireSphere/common/call/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserByPhoneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserByPhoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserByPhoneLogic {
	return &FindUserByPhoneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserByPhoneLogic) FindUserByPhone(in *user.Phone) (*user.UserInfo, error) {
	// todo: add your logic here and delete this line

	return &user.UserInfo{}, nil
}
