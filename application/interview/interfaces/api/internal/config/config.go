package config

import (
	"Ai-HireSphere/common/coze"
	"Ai-HireSphere/common/gormx"
	"Ai-HireSphere/common/redisx"
	"Ai-HireSphere/common/thrift/oss"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql gormx.Mysql
	Redis redisx.Redis
	//UserRpc zrpc.RpcClientConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	AliyunOss oss.AliyunOssConfig
	Coze      coze.Config
}
