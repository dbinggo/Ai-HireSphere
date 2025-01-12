package logic

import (
	"Ai-HireSphere/common/call/userClient"
	"context"

	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/svc"
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

func (l *FindUserByIdLogic) FindUserById(in *userClient.Id) (*userClient.UserInfo, error) {

	info, err := l.svcCtx.UserApp.FindUserById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &userClient.UserInfo{
		Id:       info.Id,
		UserName: info.UserName,
		Email:    info.Email,
		Phone:    info.Phone,
		Role:     string(info.Role),
		Sex:      string(info.Sex),
	}, nil
}
