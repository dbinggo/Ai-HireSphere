package svc

import (
	"Ai-HireSphere/application/interview/interfaces/api/internal/config"
	userClient "Ai-HireSphere/common/call/user_client"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userClient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
