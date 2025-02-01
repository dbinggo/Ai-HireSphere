package svc

import (
	"Ai-HireSphere/application/user-center/app"
	"Ai-HireSphere/application/user-center/domain/irepository"
	"Ai-HireSphere/application/user-center/infrastructure/repository"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/config"
	userClient "Ai-HireSphere/common/call/user_client"
	"Ai-HireSphere/common/gormx"
	"Ai-HireSphere/common/redisx"
	"Ai-HireSphere/common/thrift/sms"
	"Ai-HireSphere/common/zlog/dbLogger"
)

type ServiceContext struct {
	// 基本配置
	Config config.Config
	// 用户服务
	UserApp app.IUserApp
	// 基础服务
	BaseApp app.IBaseApp
	// 仓储接口
	Repo irepository.IRepoBroker
	// rpc服务
	UserRpc userClient.User
}

// 这里进行初始化各种依赖
func NewServiceContext(c config.Config) *ServiceContext {
	// 第一步先初始化数据库配置
	db := gormx.MustOpen(c.Mysql, dbLogger.New())
	rdb := redisx.MustOpen(c.Redis)
	// 第二部初始化repo仓库
	repo := repository.NewRepoBroker(db, rdb)
	// 第三部初始化rpc服务
	//userRpc := userClient.NewUser(zrpc.MustNewClient(c.UserRpc))

	// 第四部初始化APP层
	userApp := app.NewUserApp(repo, nil)

	smsClient := sms.MustNewAliyunSMSClient(c.AliyunSMS.AccessKeyId, c.AliyunSMS.AccessKeySecret)
	baseApp := app.NewBaseApp(repo, smsClient)

	return &ServiceContext{
		Config:  c,
		Repo:    repo,
		UserApp: userApp,
		BaseApp: baseApp,
	}
}
