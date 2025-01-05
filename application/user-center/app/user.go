package app

import (
	"Ai-HireSphere/application/user-center/domain/irepository"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/application/user-center/domain/services"
	"Ai-HireSphere/common/call/userClient"
	"Ai-HireSphere/common/model/enums"
	"context"
)

// 定义service
// 这里定义你对外提供的服务

type IUserApp interface {
	// 注册用户
	RegisterUser(ctx context.Context, way enums.UserRegisterWayType, user *entity.User) (string, error)
	// 查询用户
	FindUserById(ctx context.Context, id int64) (*entity.User, error)
}
type UserApp struct {
	// 这里主要是依赖
	Repo        irepository.IRepoBroker
	UserService services.IUserService
	UserRpc     userClient.User
}

func NewUserApp(repo irepository.IRepoBroker, userService services.IUserService, userRpc userClient.User) *UserApp {
	return &UserApp{
		Repo:        repo,
		UserService: userService,
		UserRpc:     userRpc,
	}
}

// RegisterUser
//
//	@Description: 注册用户 并且返回token
//	@receiver u
//	@param ctx
//	@param way
//	@param user
//	@return string
//	@return error
func (u *UserApp) RegisterUser(ctx context.Context, way enums.UserRegisterWayType, user *entity.User) (string, error) {
	// 这里是对领域服务的调用和编排
	return u.UserService.RegisterUser(user, way)
}

// FindUserById
//
//	@Description: 根据ID查询用户
//	@receiver u
//	@param ctx
//	@param id
//	@return entity.User
//	@return error
func (u *UserApp) FindUserById(ctx context.Context, id int64) (*entity.User, error) {
	return u.Repo.FindUserById(ctx, id)
}
