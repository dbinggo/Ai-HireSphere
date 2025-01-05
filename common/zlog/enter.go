package zlog

import (
	"Ai-HireSphere/common/utils"
	"Ai-HireSphere/common/zapx"
	"Ai-HireSphere/common/zlog/zeroLogger"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

func InitLogger(rest rest.RestConf) {
	path := utils.GetRootPath("")
	// zap 配置
	var zapConfig = zapx.ZapConfig{
		// 是否为 json格式
		Format: rest.Log.Encoding,
		// bug 等级
		Level: rest.Log.Level,
		// 是否开启彩色（Info 颜色）
		Colour: rest.Log.Encoding == "plain" && rest.Log.Mode == "console",
		// 日志存储路径 会在路径下生成 info.log error.log
		FilePath: rest.Log.Path,
		// 是否存储日志
		File: rest.Log.Mode == "file",
		// 是否在控制台输出
		Terminal: rest.Log.Mode == "console",
	}
	// zlog 配置
	var zlogConfig = ZlogConfig{
		// 日志格式 需要和 zapConfig.Format 一致
		Format: rest.Log.Encoding,
		// 是否开启前端日志查看
		Debug: false,
		// 是否开启caller
		Caller: true,
		// 项目路径
		Path: path,
		// 调用堆栈跳过层数 这个默认为2就行
		CallerSkip: 2,
		// 是否换行打印日志
		NewLine: rest.Log.Encoding == "plain",
		// 是否开启颜色功能
		Colour: rest.Log.Encoding == "plain" && rest.Log.Mode == "console",
		Prefix: rest.ServiceConf.Name,
	}
	logger := zapx.GetLogger(zapConfig)
	SetZlog(zlogConfig)
	initLogger(logger)
	var zapWriter logx.Writer
	if rest.Log.Encoding == "plain" {
		zapWriter = zeroLogger.NewZeroLogger(logger, rest.Log.Encoding == "json", "\n")
	} else {
		zapWriter = zeroLogger.NewZeroLogger(logger, rest.Log.Encoding == "json", "\n")
	}

	logx.SetWriter(zapWriter)
}

// 测试 生产环境专用
func Product() {
	path := utils.GetRootPath("")
	// zap 配置
	var zapConfig = zapx.ZapConfig{
		// 是否为 json格式
		Format: "json",
		// bug 等级
		Level: "info",
		// 是否开启彩色（Info 颜色）
		Colour: false,
		// 日志存储路径 会在路径下生成 info.log error.log
		FilePath: path + "/logs/",
		// 是否存储日志
		File: true,
		// 是否在控制台输出
		Terminal: true,
	}
	// zlog 配置
	var zlogConfig = ZlogConfig{
		// 日志格式 需要和 zapConfig.Format 一致
		Format: "json",
		// 是否开启前端日志查看
		Debug: false,
		// 是否开启caller
		Caller: true,
		// 项目路径
		Path: path,
		// 调用堆栈跳过层数 这个默认为2就行
		CallerSkip: 2,
		// 是否换行打印日志
		NewLine: false,
		// 是否开启颜色功能
		Colour: false,
	}
	logger := zapx.GetLogger(zapConfig)
	SetZlog(zlogConfig)
	initLogger(logger)

	zapWriter := zeroLogger.NewZeroLogger(logger, true, "")
	logx.SetWriter(zapWriter)
}

// 开发环境专用
func Develop() {
	path := utils.GetRootPath("")
	// zap 配置
	var zapConfig = zapx.ZapConfig{
		// 是否为 json格式
		Format: "plain",
		// bug 等级
		Level: "debug",
		// 是否开启彩色（Info 颜色）
		Colour: true,
		// 日志存储路径 会在路径下生成 info.log error.log
		FilePath: path + "/logs/",
		// 是否存储日志
		File: false,
		// 是否在控制台输出
		Terminal: true,
	}
	// zlog 配置
	var zlogConfig = ZlogConfig{
		// 日志格式 需要和 zapConfig.Format 一致
		Format: "plain",
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
	SetZlog(zlogConfig)
	initLogger(logger)

	zapWriter := zeroLogger.NewZeroLogger(logger, false, "\n")
	logx.SetWriter(zapWriter)
}
