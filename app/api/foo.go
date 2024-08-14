package main

import (
	"Ai-HireSphere/common/xcode"
	"Ai-HireSphere/utils"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"

	"Ai-HireSphere/app/api/internal/config"
	"Ai-HireSphere/app/api/internal/handler"
	"Ai-HireSphere/app/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", utils.GetRootPath("/app/api/etc/foo.yaml"), "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	//注册自定义响应
	httpx.SetErrorHandler(xcode.ErrHandler)
	httpx.SetOkHandler(xcode.OKHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
