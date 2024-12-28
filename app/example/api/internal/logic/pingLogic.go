package logic

import (
	"Ai-HireSphere/app/example/api/internal/svc"
	"Ai-HireSphere/common/repository"
	"Ai-HireSphere/common/zlog"
	"context"
	"github.com/zeromicro/go-zero/core/trace"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	repo   repository.RepoBroker
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		repo:   repository.Repo,
	}
}

func (l *PingLogic) Ping() error {
	// todo: add your logic here and delete this line
	logx.Info("ping")
	logx.WithContext(l.ctx).Info("ping")
	l.Logger.Infof("ping")
	zlog.InfofCtx(l.ctx, "ping")
	zlog.InfofCtx(l.ctx, "for test")
	zlog.InfofCtx(l.ctx, "for test")
	zlog.InfofCtx(l.ctx, "for test")
	zlog.InfofCtx(l.ctx, "for test")

	traceId := trace.TraceIDFromContext(l.ctx)
	spanId := trace.SpanIDFromContext(l.ctx)
	logx.Infof("traceId: %s, spanId: %s", traceId, spanId)

	return nil
}
