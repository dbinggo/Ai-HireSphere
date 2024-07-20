package test

import (
	"Ai-HireSphere/common/zapx"
	"Ai-HireSphere/common/zlog"
	"Ai-HireSphere/utils"
	"context"
	"go.uber.org/zap"
	_ "go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	path := utils.GetRootPath("")
	var zapConfig = zapx.ZapConfig{
		Format:   "terminal",
		Level:    "debug",
		Colour:   true,
		FilePath: path + "/common/zlog/test/",
		File:     true,
		Terminal: true,
	}
	var zlogConfig = zlog.ZlogConfig{
		Format:     "terminal",
		Debug:      true,
		Caller:     true,
		Path:       path,
		CallerSkip: 2,
		NewLine:    true,
		Colour:     true,
	}

	logger := zapx.GetLogger(zapConfig)
	zlog.SetZlog(zlogConfig)
	zlog.InitLogger(logger)
	zlog.SetPrefix(zlog.SetBlackColour("[test]", 42) + "\t")
	zlog.Infof("test info")
	zlog.Warnf("test warn")
	zlog.Errorf("test error")
	zlog.Debugf("test debug")
	ctx := zlog.AddField(context.Background(), zap.String("traceId", zlog.SetColour("123456", 31)))
	zlog.InfofCtx(ctx, "test info")
	zlog.WarnfCtx(ctx, "test warn")
	zlog.ErrorfCtx(ctx, "test error")
	zlog.DebugfCtx(ctx, "test debug")
}
