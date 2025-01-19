package main

import (
	"Ai-HireSphere/common/codex"
	"Ai-HireSphere/common/zlog"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"Ai-HireSphere/application/user-center/interfaces/api/internal/config"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/handler"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

//go:generate goctl api go -api *.api -dir ./  --style=go_zero --home=../../../../template
//go:generate goctl api plugin -plugin goctl-swagger="swagger -filename user.json" -api user.api -dir ../../../../docs
var configFile = flag.String("f", "etc/user.yaml", "the config file")

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
