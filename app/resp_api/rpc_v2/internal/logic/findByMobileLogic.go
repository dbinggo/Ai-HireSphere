package logic

import (
	"context"

	"Ai-HireSphere/app/resp_api/rpc_v2/internal/svc"
	"Ai-HireSphere/app/resp_api/rpc_v2/service"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByMobileLogic) FindByMobile(in *service.FindByMobileRequest) (*service.FindByMobileResponse, error) {
	// todo: add your logic here and delete this line

	return &service.FindByMobileResponse{}, nil
}
