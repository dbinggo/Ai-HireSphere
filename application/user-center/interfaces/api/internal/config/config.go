package config

import (
	"Ai-HireSphere/common/gormx"
	"Ai-HireSphere/common/redisx"
	"Ai-HireSphere/common/thrift/sms"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

// 这里进行各种自定义配置
type Config struct {
	rest.RestConf
	Mysql   gormx.Mysql
	Redis   redisx.Redis
	UserRpc zrpc.RpcClientConf
	Auth    struct {
		AccessSecret string
		AccessExpire int64
	}
	AliyunSMS sms.AliyunSMSConfig
}
