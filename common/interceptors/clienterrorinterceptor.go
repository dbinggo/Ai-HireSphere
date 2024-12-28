package interceptors

import (
	"context"

	"Ai-HireSphere/common/xcode"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

/*
rpc client拦截器
使用说明：
在api层的svc中new rpc服务时，加入这个拦截器即可，示例如下：
func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	return &ServiceContext{
		Config:   c,
		UserRPC:  user.NewUser(userRPC),
	}
}

*/

func ClientErrorInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		err := invoker(ctx, method, req, reply, cc, opts...)
		if err != nil {
			grpcStatus, _ := status.FromError(err)
			xc := xcode.GrpcStatusToXCode(grpcStatus)
			err = errors.WithMessage(xc, grpcStatus.Message())
		}

		return err
	}
}
