package services

import (
	"Ai-HireSphere/application/user-center/domain/irepository"
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/common/call/userClient"
	"Ai-HireSphere/common/model/enums"
	"context"
)

// 本层为用户服务层
// 主要为编排实体类 领域服务 领域模型 聚合根 仓储 等

// IUserService
// @Description: 对外提供的接口
type IUserService interface {
	RegisterUser(user *entity.User, way enums.UserRegisterWayType) (token string, err error)
	LoginUser(user *entity.User, way enums.UserRegisterWayType) (token string, err error)
}

type UserService struct {
	ctx     context.Context
	useRepo idataaccess.IUserGorm
	userRpc userClient.User
}

func NewUserService(repo irepository.IRepoBroker, userRpc userClient.User) IUserService {
	return &UserService{
		useRepo: repo,
		userRpc: userRpc,
	}
}

func (s *UserService) RegisterUser(user *entity.User, way enums.UserRegisterWayType) (token string, err gerr.Error) {
	// 调用领域模型方法进行注册
	if err = user.Register(way); err != nil {
		return token, err
	}

	// 调用仓储方法进行保存
	user, err = s.useRepo.SaveUser(s.ctx, user)
	if err != nil {
		return token, err
	}

	// 生成token
	token, err = user.GenerateToken()
	return token, err
}

func (s *UserService) LoginUser(user *entity.User, way enums.UserRegisterWayType) (token string, err error) {
	// 调用仓储查找这个user
	data := ""
	switch way {
	case enums.UserRegisterWayTypeEmail:
		data = user.Email
	case enums.UserRegisterWayTypePhone:
		data = user.Phone
	}
	user, err = s.useRepo.FindUserByLoginType(s.ctx, way, data)
	if err != nil {
		return token, err
	}
	return user.GenerateToken()

}
