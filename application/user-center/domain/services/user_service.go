package services

import (
	"Ai-HireSphere/application/user-center/domain/irepository"
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	userClient "Ai-HireSphere/common/call/userrpc"
	"Ai-HireSphere/common/codex"
	"Ai-HireSphere/common/model/enums"
	"context"
	"github.com/dbinggo/gerr"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// 本层为用户服务层
// 主要为编排实体类 领域服务 领域模型 聚合根 仓储 等

// IUserService
// @Description: 对外提供的接口
type IUserService interface {
	RegisterUser(user *entity.UserEntity, way enums.UserRegisterMethodType, data string) (token string, err gerr.Error)
	LoginUser(user *entity.UserEntity, way enums.UserRegisterMethodType) (token string, err gerr.Error)
}

type UserService struct {
	ctx     context.Context
	useRepo idataaccess.IUserGorm
	userRpc userClient.UserRpc
}

func NewUserService(repo irepository.IRepoBroker, userRpc userClient.UserRpc) IUserService {
	return &UserService{
		useRepo: repo,
		userRpc: userRpc,
	}
}

func (s *UserService) RegisterUser(user *entity.UserEntity, way enums.UserRegisterMethodType, data string) (token string, err gerr.Error) {

	// 调用领域模型方法进行注册
	if err = user.Register(way); err != nil {
		return token, err
	}
	// 查看是否有这个用户
	_, err = s.useRepo.FindUserByLoginType(s.ctx, way, data)
	if err == nil {
		return token, gerr.WithStack(codex.UserRegisterExist)
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return token, gerr.Wraps(codex.ServerErr, err)
	}

	// 调用仓储方法进行保存
	*user, err = s.useRepo.SaveUser(s.ctx, *user)
	if err != nil {
		return token, err
	}

	// 进阶保存方法
	token, err = user.GenerateToken()
	return token, err
}

func (s *UserService) LoginUser(user *entity.UserEntity, way enums.UserRegisterMethodType) (token string, err gerr.Error) {
	// 调用仓储查找这个user
	data := ""
	switch way {
	case enums.UserRegisterWayTypeEmail:
		data = user.Email
	case enums.UserRegisterWayTypePhone:
		data = user.Phone
	}
	*user, err = s.useRepo.FindUserByLoginType(s.ctx, way, data)
	if err != nil {
		return token, err
	}
	return user.GenerateToken()
}
