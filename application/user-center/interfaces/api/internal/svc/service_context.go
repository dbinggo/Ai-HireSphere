package svc

import (
	"Ai-HireSphere/application/user-center/app"
	"Ai-HireSphere/application/user-center/domain/irepository"
	"Ai-HireSphere/application/user-center/infrastructure/repository"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/config"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/middleware"
	userClient "Ai-HireSphere/common/call/userrpc"
	"Ai-HireSphere/common/gormx"
	"Ai-HireSphere/common/redisx"
	"Ai-HireSphere/common/thrift/sms"
	"Ai-HireSphere/common/zlog/dbLogger"
	"github.com/zeromicro/go-zero/rest"
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
	UserRpc userClient.UserRpc

	CorsMiddleware rest.Middleware
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

	smsClient := sms.MustNewAliyunSMSClient(c.AliyunSMS.AccessKeyId, c.AliyunSMS.AccessKeySecret)
	userApp := app.NewUserApp(repo, nil, smsClient)
	baseApp := app.NewBaseApp(repo, smsClient)

	return &ServiceContext{
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
		Config:         c,
		Repo:           repo,
		UserApp:        userApp,
		BaseApp:        baseApp,
	}
}
