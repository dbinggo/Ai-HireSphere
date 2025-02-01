package user

import (
	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/types"
	"Ai-HireSphere/common/utils"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfo, err error) {

	userId := utils.GetUserId(l.ctx)
	user, err := l.svcCtx.UserApp.FindUserById(l.ctx, userId)
	if err != nil {
		return nil, err
	}
	return &types.UserInfo{
		Id:       user.Id,
		Username: user.UserName,
		Avatar:   user.Avatar,
		Role:     string(user.Role),
		Phone:    user.Phone,
		Email:    user.Phone,
		Sex:      int(user.Sex),
	}, nil
}
