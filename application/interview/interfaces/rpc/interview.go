package main

import (
	"Ai-HireSphere/common/interceptors"
	"Ai-HireSphere/common/zlog"
	"flag"
	"fmt"

	"Ai-HireSphere/application/interview/interfaces/rpc/internal/config"
	"Ai-HireSphere/application/interview/interfaces/rpc/internal/server"
	"Ai-HireSphere/application/interview/interfaces/rpc/internal/svc"
	"Ai-HireSphere/common/call/interview"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/interview.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		interview.RegisterInterviewRpcServer(grpcServer, server.NewInterviewRpcServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})

	// 自定义拦截器
	s.AddUnaryInterceptors(interceptors.ServerErrorInterceptor())

	defer s.Stop()
	//注册自定义日志
	zlog.InitLogger(c.ServiceConf)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
