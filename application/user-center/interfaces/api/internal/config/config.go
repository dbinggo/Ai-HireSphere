package config

import (
	"Ai-HireSphere/common/gormx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// 这里进行各种自定义配置
type Config struct {
	rest.RestConf
	Mysql gormx.Mysql

	UserRpc zrpc.RpcClientConf
}
