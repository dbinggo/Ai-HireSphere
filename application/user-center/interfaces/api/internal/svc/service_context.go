package svc

import (
	"Ai-HireSphere/application/user-center/app"
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	"Ai-HireSphere/application/user-center/domain/services"
	"Ai-HireSphere/application/user-center/infrastructure/repository"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/config"
	"Ai-HireSphere/common/call/userClient"
	"Ai-HireSphere/common/gormx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserApp app.IUserApp
	Repo    idataaccess.IUserGorm
}

// 这里进行初始化各种依赖
func NewServiceContext(c config.Config) *ServiceContext {
	// 第一步先初始化数据库配置
	db := gormx.MustOpen(c.Mysql, nil)

	// 第二部初始化repo仓库
	repo := repository.NewRepoBroker(db, nil)

	// 第三部初始化rpc服务
	userRpc := userClient.NewUser(zrpc.MustNewClient(c.UserRpc))

	// 第四步初始化领域层服务
	userService := services.NewUserService(repo, userRpc)

	// 第五部初始化APP层
	userApp := app.NewUserApp(repo, userService, userRpc)

	return &ServiceContext{
		Config:  c,
		Repo:    repo,
		UserApp: userApp,
	}
}
