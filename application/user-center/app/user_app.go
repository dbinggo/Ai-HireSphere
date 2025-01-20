package app

import (
	"Ai-HireSphere/application/user-center/domain/irepository"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/application/user-center/domain/services"
	userClient "Ai-HireSphere/common/call/user_client"
	"Ai-HireSphere/common/model/enums"
	"context"
	"github.com/dbinggo/gerr"
)

// 定义service
// 这里定义你对外提供的服务

type IUserApp interface {
	// 注册用户
	RegisterUser(ctx context.Context, way enums.UserRegisterWayType, data string, code string) (string, gerr.Error)
	// 查询用户
	FindUserById(ctx context.Context, id int64) (*entity.User, gerr.Error)
	// 用户登陆
	LoginUser(ctx context.Context, way enums.UserRegisterWayType, data string, code string) (string, gerr.Error)
	// 登录用户
}
type UserApp struct {
	// 这里主要是依赖
	Repo    irepository.IRepoBroker
	UserRpc userClient.User
}

func NewUserApp(repo irepository.IRepoBroker, userRpc userClient.User) *UserApp {
	return &UserApp{
		Repo:    repo,
		UserRpc: userRpc,
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
func (u *UserApp) RegisterUser(ctx context.Context, way enums.UserRegisterWayType, data string, code string) (string, gerr.Error) {
	// 这里是对领域服务的调用和编排
	// 首先校验code正确性

	// 基础校验校验验证码
	if err := services.NewBaseCaptcha(ctx, u.Repo, nil).CaptchaCheck(enums.CaptchaWayTypeRegister, data, code); err != nil {
		return "", err
	}
	// 用户服务注册
	s := services.NewUserService(u.Repo, u.UserRpc)

	// 创建一个user
	user := &entity.User{
		Role: enums.UserRoleTypeUser,
	}

	switch way {
	case enums.UserRegisterWayTypeEmail:
		user.Email = data
	case enums.UserRegisterWayTypePhone:
		user.Phone = data
	}

	// 用户注册服务
	return s.RegisterUser(user, way)
}

// FindUserById
//
//	@Description: 根据ID查询用户
//	@receiver u
//	@param ctx
//	@param id
//	@return entity.User
//	@return error
func (u *UserApp) FindUserById(ctx context.Context, id int64) (*entity.User, gerr.Error) {
	return u.Repo.FindUserById(ctx, id)
}

func (u *UserApp) LoginUser(ctx context.Context, way enums.UserRegisterWayType, data string, code string) (string, gerr.Error) {

	// 基础校验校验验证码
	if err := services.NewBaseCaptcha(ctx, u.Repo, nil).CaptchaCheck(enums.CaptchaWayTypeLogin, data, code); err != nil {
		return "", err
	}

	user := &entity.User{}
	switch way {
	case enums.UserRegisterWayTypeEmail:
		user.Email = data
	case enums.UserRegisterWayTypePhone:
		user.Phone = data

	}

	return services.NewUserService(u.Repo, u.UserRpc).LoginUser(user, way)
}
