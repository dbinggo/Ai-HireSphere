package svc

import (
	"Ai-HireSphere/application/user-center/app"
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	"Ai-HireSphere/application/user-center/domain/services"
	"Ai-HireSphere/application/user-center/infrastructure/repository"
	"Ai-HireSphere/application/user-center/interfaces/rpc/internal/config"
	"Ai-HireSphere/common/gormx"
)

type ServiceContext struct {
	Config  config.Config
	UserApp app.IUserApp
	Repo    idataaccess.IUserGorm
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 第一步先初始化数据库配置
	db := gormx.MustOpen(c.Mysql, nil)

	// 第二部初始化repo仓库
	repo := repository.NewRepoBroker(db, nil)

	// 第三步初始化领域层服务
	userService := services.NewUserService(repo, nil)

	// 第四部初始化APP层
	userApp := app.NewUserApp(repo, userService, nil)

	return &ServiceContext{
		Config:  c,
		UserApp: userApp,
		Repo:    repo,
	}
}
