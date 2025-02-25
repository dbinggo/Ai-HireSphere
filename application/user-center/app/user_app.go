package app

import (
	"Ai-HireSphere/application/user-center/domain/irepository"
	"Ai-HireSphere/application/user-center/domain/irepository/isms"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/application/user-center/domain/services"
	userClient "Ai-HireSphere/common/call/userrpc"
	"Ai-HireSphere/common/model/enums"
	"context"
	"github.com/dbinggo/gerr"
)

// 定义service
// 这里定义你对外提供的服务

type IUserApp interface {
	// 注册用户
	RegisterUser(ctx context.Context, method enums.UserRegisterMethodType, data string, code string) (string, gerr.Error)
	// 查询用户
	FindUserById(ctx context.Context, id int64) (entity.UserEntity, gerr.Error)
	// 用户登陆
	LoginUser(ctx context.Context, method enums.UserRegisterMethodType, data string, code string) (string, gerr.Error)
}
type UserApp struct {
	// 这里主要是依赖
	Repo    irepository.IRepoBroker
	UserRpc userClient.UserRpc
	Sms     isms.ISms
}

func NewUserApp(repo irepository.IRepoBroker, userRpc userClient.UserRpc, sms isms.ISms) *UserApp {
	return &UserApp{
		Repo:    repo,
		UserRpc: userRpc,
		Sms:     sms,
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
func (u *UserApp) RegisterUser(ctx context.Context, way enums.UserRegisterMethodType, data string, code string) (string, gerr.Error) {
	// 这里是对领域服务的调用和编排
	// 首先校验code正确性

	// 基础校验校验验证码
	if err := services.NewBaseCaptcha(ctx, u.Repo, u.Sms).CaptchaCheck(enums.CaptchaWayTypeRegister, data, code); err != nil {
		return "", err
	}
	// 用户服务注册
	s := services.NewUserService(u.Repo, u.UserRpc)

	// 创建一个user
	user := &entity.UserEntity{
		Role: enums.UserRoleTypeUser,
	}

	switch way {
	case enums.UserRegisterWayTypeEmail:
		user.Email = data
	case enums.UserRegisterWayTypePhone:
		user.Phone = data
	}

	// 用户注册服务
	return s.RegisterUser(user, way, data)
}

// FindUserById
//
//	@Description: 根据ID查询用户
//	@receiver u
//	@param ctx
//	@param id
//	@return entity.UserEntity
//	@return error
func (u *UserApp) FindUserById(ctx context.Context, id int64) (entity.UserEntity, gerr.Error) {
	return u.Repo.FindUserById(ctx, id)
}

func (u *UserApp) LoginUser(ctx context.Context, method enums.UserRegisterMethodType, data string, code string) (string, gerr.Error) {

	// 基础校验校验验证码
	if err := services.NewBaseCaptcha(ctx, u.Repo, u.Sms).CaptchaCheck(enums.CaptchaWayTypeLogin, data, code); err != nil {
		return "", err
	}

	user := &entity.UserEntity{}
	switch method {
	case enums.UserRegisterWayTypeEmail:
		user.Email = data
	case enums.UserRegisterWayTypePhone:
		user.Phone = data

	}

	return services.NewUserService(u.Repo, u.UserRpc).LoginUser(user, method)
}
