package svc

import (
	"Ai-HireSphere/app/example/api/internal/config"
)

type ServiceContext struct {
	Config config.Config
}

func NewServiceContext(c config.Config) *ServiceContext {
	rpc = Newrpc(options...)
	return &ServiceContext{
		Config: c,
		BD:     db,
	}
}
