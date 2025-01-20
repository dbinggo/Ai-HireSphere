package logic

import (
	"context"

	"Ai-HireSphere/application/interview/interfaces/rpc/internal/svc"
	"Ai-HireSphere/common/call/interview"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *interview.Request) (*interview.Response, error) {
	// todo: add your logic here and delete this line

	return &interview.Response{}, nil
}
