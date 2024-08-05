package logic

import (
	"Ai-HireSphere/app/resp_api/api/internal/svc"
	"Ai-HireSphere/app/resp_api/api/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PingLogic) Ping() (resp *types.Resp, err error) {
	resp = new(types.Resp)
	// todo: add your logic here and delete this line
	resp.Data = "123"
	return resp, nil
}
