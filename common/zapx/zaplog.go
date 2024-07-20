package zapx

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

/*
[感谢伟人 让我彻底搞懂 zap]https://juejin.cn/post/7313979344561242162?searchId=20240613163846377BACC6CC0FB80CC369
*/

type ZapConfig struct {
	Format   string
	Level    string
	Colour   bool
	FilePath string
	File     bool
	Terminal bool
}

func GetLogger(config ZapConfig) *zap.Logger {
	var coreConfigs = make([]zapConfig, 0)
	var cors = make([]zapcore.Core, 0)

	var encoder func(cfg zapcore.EncoderConfig) zapcore.Encoder
	switch config.Format {
	case "json":
		encoder = zapcore.NewJSONEncoder
	default:
		encoder = zapcore.NewConsoleEncoder
	}
	var needColour = false
	// 彩色使用位置 非json 且为开发模式
	if config.Colour {
		needColour = true
	}
	level := transLevel(config.Level)

	if config.File {
		//本开发模式旨在将正常信息及以上的log记录在文件中，方便查看
		fileInfoConfig := newZapConfig().
			setEncoder(needColour, encoder).
			setFileWriteSyncer(config.FilePath + "info.log").
			setLevelEnabler(level).
			getConfig()
		fileInfoCore := fileInfoConfig.getCore()
		//本开发模式旨在将error及以上的log记录在文件中，方便查看
		fileErrorConfig := newZapConfig().
			setEncoder(needColour, encoder).
			setFileWriteSyncer(config.FilePath + "error.log").
			setLevelEnabler(zapcore.ErrorLevel).
			getConfig()
		fileErrorCore := fileErrorConfig.getCore()
		coreConfigs = append(coreConfigs, fileInfoConfig, fileErrorConfig)
		cors = append(cors, fileInfoCore, fileErrorCore)
	}
	if config.Terminal {
		//输出在控制台
		consoleInfoConfig := newZapConfig().
			setEncoder(needColour, encoder).
			setStdOutWriteSyncer().
			setLevelEnabler(level).
			getConfig()
		consoleInfoCore := consoleInfoConfig.getCore()
		coreConfigs = append(coreConfigs, consoleInfoConfig)
		cors = append(cors, consoleInfoCore)
	}
	zapLogger := makeZapLogger(cors, zap.AddCallerSkip(1))
	defer zapLogger.Sync()
	return zapLogger
}
func makeZapLogger(cors []zapcore.Core, options ...zap.Option) *zap.Logger {
	core := zapcore.NewTee(cors...)
	return zap.New(core, options...)

}

type zapConfig struct {
	core             zapcore.Core
	encoder          zapcore.Encoder
	writeSyncerSlice []zapcore.WriteSyncer
	levelEnabler     zapcore.LevelEnabler
}

func newZapConfig() *zapConfig {
	return &zapConfig{
		writeSyncerSlice: make([]zapcore.WriteSyncer, 0),
	}
}
func (z *zapConfig) getConfig() zapConfig {
	return *z
}

// 定制core
func (z *zapConfig) getCore() zapcore.Core {
	z.core = zapcore.NewCore(z.encoder, zapcore.NewMultiWriteSyncer(z.writeSyncerSlice...), z.levelEnabler)
	return z.core
}

// encoder 是编码器，以什么样的格式写入日志。
// 目前，zap只支持两种编码器——JSON Encoder和Console Encoder
// 储存在日志中的文件就不要颜色了
func (z *zapConfig) setEncoder(needColour bool, encoder func(cfg zapcore.EncoderConfig) zapcore.Encoder) *zapConfig {
	encodeLevel := zapcore.CapitalLevelEncoder
	if needColour {
		encodeLevel = zapcore.CapitalColorLevelEncoder
	}
	z.encoder = encoder(zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "name",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    encodeLevel,
		EncodeTime:     newTimeEncoder(),
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	})

	return z
}
func (z zapConfig) getEncoder() zapcore.Encoder {
	return z.encoder
}
func (z *zapConfig) setFileWriteSyncer(logFilePath string) *zapConfig {
	//引入第三方库 Lumberjack 加入日志切割功能
	lumberWriteSyncer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    1024,  // megabytes
		MaxBackups: 7,     //最多备份文件数量
		MaxAge:     28,    // days
		Compress:   false, //Compress确定是否应该使用gzip压缩已旋转的日志文件。默认值是不执行压缩。
	}
	z.writeSyncerSlice = append(z.writeSyncerSlice, zapcore.AddSync(lumberWriteSyncer))

	return z
}
func (z zapConfig) getWriteSyncers() []zapcore.WriteSyncer {
	return z.writeSyncerSlice
}
func (z *zapConfig) setStdOutWriteSyncer() *zapConfig {
	z.writeSyncerSlice = append(z.writeSyncerSlice, zapcore.AddSync(os.Stdout))
	return z
}
func (z zapConfig) getLevelEnabler() zapcore.LevelEnabler {
	return z.levelEnabler
}

func (z *zapConfig) setLevelEnabler(enabler zapcore.Level) *zapConfig {
	z.levelEnabler = zap.LevelEnablerFunc(func(lev zapcore.Level) bool { //error级别
		return lev >= enabler
	})
	return z
}

func newTimeEncoder() zapcore.TimeEncoder {
	return func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006/1/2 15:04:05.000"))
	}
}

func transLevel(lev string) zapcore.Level {
	var levels = []string{
		"debug",
		"info",
		"warn",
		"error",
		"panic",
	}
	for power, level := range levels {
		if level == lev {
			return zapcore.Level(power - 1)
		}
	}
	return zapcore.DebugLevel
}
