package test

import (
	"Ai-HireSphere/common/zapx"
	"Ai-HireSphere/common/zlog"
	"Ai-HireSphere/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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
	ctx := zlog.SetPrefix(context.Background(), zlog.SetBlackColour("[test]", 42))
	zlog.Infof("test info")
	zlog.Warnf("test warn")
	zlog.Errorf("test error")
	zlog.Debugf("test debug")
	ctx = zlog.AddField(ctx, zap.String("traceId", zlog.SetColour("123456", 31)))
	zlog.InfofCtx(ctx, "test info")
	zlog.WarnfCtx(ctx, "test warn")
	zlog.ErrorfCtx(ctx, "test error")
	zlog.DebugfCtx(ctx, "test debug")
}
func TestZapX(t *testing.T) {
	path := utils.GetRootPath("")
	// zap 配置
	var zapConfig = zapx.ZapConfig{
		// 是否为 json格式
		Format: "terminal",
		// bug 等级
		Level: "debug",
		// 是否开启彩色（Info 颜色）
		Colour: true,
		// 日志存储路径 会在路径下生成 info.log error.log
		FilePath: path + "/common/zlog/test/",
		// 是否存储日志
		File: true,
		// 是否在控制台输出
		Terminal: true,
	}

	// zlog 配置
	var zlogConfig = zlog.ZlogConfig{
		// 日志格式 需要和 zapConfig.Format 一致
		Format: "terminal",
		// 是否开启前端日志查看
		Debug: true,
		// 是否开启caller
		Caller: true,
		// 项目路径
		Path: path,
		// 调用堆栈跳过层数 这个默认为2就行
		CallerSkip: 2,
		// 是否换行打印日志
		NewLine: true,
		// 是否开启颜色功能
		Colour: true,
	}
	logger := zapx.GetLogger(zapConfig)
	zlog.SetZlog(zlogConfig)
	zlog.InitLogger(logger)

	zapWriter := zlog.NewZapWriter(logger)
	logx.SetWriter(zapWriter)
	ctx := zlog.AddField(context.Background(), zap.String("traceId", zlog.SetColour("123456", 31)), zap.String("spanId", zlog.SetColour("123456", 31)))

	logx.Infof("test info")

	zlog.InfofCtx(ctx, "test info")

	logx.Stat("test stat")
}
