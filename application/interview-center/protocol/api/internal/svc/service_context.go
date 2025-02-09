package svc

import (
	"Ai-HireSphere/application/interview-center/entity"
	"Ai-HireSphere/application/interview-center/protocol/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Entity *entity.Entity
}

func NewServiceContext(c config.Config) *ServiceContext {
	e := entity.New()
	return &ServiceContext{
		Config: c,
		Entity: e,
	}
}
