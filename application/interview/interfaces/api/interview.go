package main

import (
	"Ai-HireSphere/common/codex"
	"Ai-HireSphere/common/zlog"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"Ai-HireSphere/application/interview/interfaces/api/internal/config"
	"Ai-HireSphere/application/interview/interfaces/api/internal/handler"
	"Ai-HireSphere/application/interview/interfaces/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/interview-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册自定义响应
	httpx.SetErrorHandler(codex.ErrHandler)
	httpx.SetOkHandler(codex.OKHandler)
	// 注册自定义日志
	zlog.InitLogger(c.ServiceConf)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
