package svc

import (
	repot "Ai-HireSphere/application/interview/infrastructure/repository"
	"Ai-HireSphere/application/interview/interfaces/api/internal/config"
	"Ai-HireSphere/application/interview/interfaces/api/internal/middleware"
	userClient "Ai-HireSphere/common/call/userrpc"
	"Ai-HireSphere/common/gormx"
	"Ai-HireSphere/common/thrift/oss"
	"Ai-HireSphere/common/zlog/dbLogger"
	"github.com/zeromicro/go-zero/rest"

	"Ai-HireSphere/application/interview/app"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userClient.UserRpc
	Auth    struct {
		AccessSecret string
		AccessExpire int64
	}
	DeepSeek struct {
		ApiKey string
	}
	ResumeAPP      app.IResumeApp
	CorsMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 第一步初始化oss服务
	client := oss.MustNewAliyunOSSClient(c.AliyunOss)
	// 第二部初始化repo服务
	db := gormx.MustOpen(c.Mysql, dbLogger.New())
	repo := repot.NewRepoBroker(db, nil)

	resumeApp := app.NewResumeApp(client, repo)

	return &ServiceContext{
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,

		Config:    c,
		ResumeAPP: resumeApp,
	}
}
