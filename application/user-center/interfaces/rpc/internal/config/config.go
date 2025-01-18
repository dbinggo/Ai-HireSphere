package config

import (
	"Ai-HireSphere/common/gormx"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf

	Mysql gormx.Mysql
}
