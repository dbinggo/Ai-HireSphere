package svc

import (
	repot "Ai-HireSphere/application/interview/infrastructure/repository"
	"Ai-HireSphere/application/interview/interfaces/api/internal/config"
	userClient "Ai-HireSphere/common/call/user_client"
	"Ai-HireSphere/common/gormx"
	"Ai-HireSphere/common/thrift/oss"
	"Ai-HireSphere/common/zlog/dbLogger"

	"Ai-HireSphere/application/interview/app"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userClient.User
	Auth    struct {
		AccessSecret string
		AccessExpire int64
	}
	ResumeAPP app.IResumeApp
}

func NewServiceContext(c config.Config) *ServiceContext {

	// 第一步初始化oss服务
	client := oss.MustNewAliyunOSSClient(c.AliyunOss)
	// 第二部初始化repo服务
	db := gormx.MustOpen(c.Mysql, dbLogger.New())
	repo := repot.NewRepoBroker(db, nil)

	resumeApp := app.NewResumeApp(client, repo)

	return &ServiceContext{
		Config:    c,
		ResumeAPP: resumeApp,
	}
}
