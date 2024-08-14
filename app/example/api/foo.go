package main

import (
	"Ai-HireSphere/app/example/api/internal/config"
	"Ai-HireSphere/app/example/api/internal/handler"
	"Ai-HireSphere/app/example/api/internal/svc"
	"Ai-HireSphere/common/zlog"
	"Ai-HireSphere/utils"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", utils.GetRootPath("/app/example/etc/foo.yaml"), "the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	zlog.Develop()
	logx.Infof("Starting server at %s:%d...", c.Host, c.Port)
	server.Start()
}
