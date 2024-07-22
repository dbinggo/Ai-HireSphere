package zlog

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
)

const (
	logger_formate_json = "json"

	loggerDebugKey  = "debug"
	loggerCallerKey = "caller"
	loggerKey       = "logger"
	loggerExKey     = "loggerEx"
	loggerPrefixKey = "loggerPrefixKey"
	// 分隔符
	formatSeparator = "%v\t"
)

type ZlogConfig struct {
	Format     string
	Debug      bool
	Caller     bool
	Path       string
	CallerSkip int
	NewLine    bool
	Colour     bool
}

var logger *zap.Logger
var zlogConfig *ZlogConfig
var newLine = "\n"

func SetZlog(config ZlogConfig) {
	zlogConfig = &config
	if !config.NewLine {
		newLine = ""
	}
}
func SetPrefix(ctx context.Context, prefix string) context.Context {

	ctx = context.WithValue(ctx, loggerPrefixKey, prefix)

	return ctx
}

// SetColour 设置字体颜色 30黑 31红 32绿 33黄 34蓝 35紫 36青 37白
func SetColour(text string, colour int) string {
	if !needColour() {
		return text
	}
	return fmt.Sprintf("\u001b[%dm%s\u001b[0m", colour, text)

}

// SetBlackColour 设置背景颜色 40黑 41红 42绿 43黄 44蓝 45紫 46青 47白
func SetBlackColour(text string, colour int) string {
	if !needColour() {
		return text
	}
	return fmt.Sprintf("\u001b[%dm%s\u001b[0m", colour, text)

}
func formatJson() bool {
	return zlogConfig.Format == logger_formate_json
}
func debug() bool {

	return zlogConfig.Debug
}
func needColour() bool {
	return zlogConfig.Colour
}
func needCaller() bool {

	return zlogConfig.Caller
}

// AddField
//
//	@Description:给指定context添加字段 实现类似traceid作用
//	@param ctx
//	@param fields
//	@return context.Context
func AddField(ctx context.Context, fields ...zapcore.Field) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	if formatJson() {
		ctx = context.WithValue(ctx, loggerKey, withContext(ctx).With(fields...))
	} else {
		ctx = context.WithValue(ctx, loggerKey, withContext(ctx))
		if ex, ok := ctx.Value(loggerExKey).([]zapcore.Field); ok {
			ex = append(ex, fields...)
			context.WithValue(ctx, loggerExKey, ex)
		} else {
			ctx = context.WithValue(ctx, loggerExKey, fields)
		}
	}
	if debug() {
		if _, ok := ctx.Value(loggerDebugKey).([]string); !ok {
			context.WithValue(ctx, loggerDebugKey, []string{})
		}

	}
	return ctx
}
func InitLogger(zapLogger *zap.Logger) {
	logger = zapLogger
}

// 从指定的context返回一个zap实例
func withContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return logger
	}
	if ctxLogger, ok := ctx.Value(loggerKey).(*zap.Logger); ok {
		return ctxLogger
	}
	return logger
}

func Infof(format string, v ...interface{}) {
	_logger, formatCaller, vCaller := addCaller(logger)
	if !formatJson() {
		v = append(vCaller, v...)
		_logger.Info(fmt.Sprintf(formatCaller+newLine+format, v...))

	} else {
		_logger.Info(fmt.Sprintf(format, v...))
	}
}

func Errorf(format string, v ...interface{}) {
	_logger, formatCaller, vCaller := addCaller(logger)
	v = append(vCaller, v...)

	_logger.Error(fmt.Sprintf(formatCaller+newLine+format, v...))
}

func Warnf(format string, v ...interface{}) {
	_logger, formatCaller, vCaller := addCaller(logger)
	v = append(vCaller, v...)
	_logger.Warn(fmt.Sprintf(formatCaller+newLine+format, v...))
}

func Debugf(format string, v ...interface{}) {
	_logger, formatCaller, vCaller := addCaller(logger)
	v = append(vCaller, v...)
	_logger.Debug(fmt.Sprintf(formatCaller+newLine+format, v...))
}

func Panicf(format string, v ...interface{}) {
	_logger, formatCaller, vCaller := addCaller(logger)
	v = append(vCaller, v...)
	_logger.Panic(fmt.Sprintf(formatCaller+newLine+format, v...))
}

func Fatalf(format string, v ...interface{}) {
	_logger, formatCaller, vCaller := addCaller(logger)
	v = append(vCaller, v...)
	_logger.Fatal(fmt.Sprintf(formatCaller+newLine+format, v...))
}

func addExField(ctx context.Context) (string, []interface{}) {
	var formatPre string
	v := make([]interface{}, 0)
	if !formatJson() {
		if exField, ok := ctx.Value(loggerExKey).([]zap.Field); ok {
			for _, field := range exField {
				v = append(v, field.String)
				formatPre = formatPre + formatSeparator
			}
		}
	}
	return formatPre, v
}
func addCaller(_logger *zap.Logger) (zap.Logger, string, []interface{}) {
	if !needCaller() {
		return *_logger, "", nil
	}
	format := "%s:%d"
	_, file, line, _ := runtime.Caller(zlogConfig.CallerSkip)
	_v := make([]interface{}, 0)
	file = file[len(zlogConfig.Path)+1:]
	_v = append(_v, file, line)
	if formatJson() {
		_logger = _logger.With(zap.String(loggerCallerKey, fmt.Sprintf(format, file, line)))
		return *_logger, "", []interface{}{}
	}
	return *logger, format + "\t", _v
}
func addDebugMessage(ctx context.Context, message string) {
	if debug() {
		if log, ok := ctx.Value(loggerDebugKey).(*[]string); ok {
			*log = append(*log, message)
			go func() {}()
		} else {
			ctx = context.WithValue(ctx, loggerDebugKey, []string{message})
		}
	}
}

// 下面的logger方法会携带trace id

func InfofCtx(ctx context.Context, format string, v ...interface{}) {
	formatField, vField := addExField(ctx)
	_logger, formatCaller, vCaller := addCaller(withContext(ctx))
	addDebugMessage(ctx, fmt.Sprintf("[INFO]  "+formatCaller, vCaller...))
	addDebugMessage(ctx, fmt.Sprintf(format, v...))
	_v := append(vField, v...)
	v = append(vCaller, _v...)
	if prefix, ok := ctx.Value(loggerPrefixKey).(string); ok {
		format = prefix + "\t" + format
	}
	_logger.Info(fmt.Sprintf(formatCaller+formatField+newLine+format, v...))
}

func ErrorfCtx(ctx context.Context, format string, v ...interface{}) {
	formatField, vField := addExField(ctx)
	_logger, formatCaller, vCaller := addCaller(withContext(ctx))
	addDebugMessage(ctx, fmt.Sprintf("[ERROR] "+formatCaller, vCaller...))
	addDebugMessage(ctx, fmt.Sprintf(format, v...))
	_v := append(vField, v...)
	v = append(vCaller, _v...)
	if prefix, ok := ctx.Value(loggerPrefixKey).(string); ok {
		format = prefix + "\t" + format
	}
	_logger.Error(fmt.Sprintf(formatCaller+formatField+newLine+format, v...))
}

func WarnfCtx(ctx context.Context, format string, v ...interface{}) {
	formatField, vField := addExField(ctx)
	_logger, formatCaller, vCaller := addCaller(withContext(ctx))
	addDebugMessage(ctx, fmt.Sprintf("[WARN]  "+formatCaller, vCaller...))
	addDebugMessage(ctx, fmt.Sprintf(format, v...))
	_v := append(vField, v...)
	v = append(vCaller, _v...)
	if prefix, ok := ctx.Value(loggerPrefixKey).(string); ok {
		format = prefix + "\t" + format
	}
	_logger.Warn(fmt.Sprintf(formatCaller+formatField+newLine+format, v...))
}

func DebugfCtx(ctx context.Context, format string, v ...interface{}) {
	formatField, vField := addExField(ctx)
	_logger, formatCaller, vCaller := addCaller(withContext(ctx))
	addDebugMessage(ctx, fmt.Sprintf("[DEBUG] "+formatCaller, vCaller...))
	addDebugMessage(ctx, fmt.Sprintf(format, v...))
	_v := append(vField, v...)
	v = append(vCaller, _v...)
	if prefix, ok := ctx.Value(loggerPrefixKey).(string); ok {
		format = prefix + format
	}
	_logger.Debug(fmt.Sprintf(formatCaller+formatField+newLine+format, v...))
}

func PanicfCtx(ctx context.Context, format string, v ...interface{}) {
	formatField, vField := addExField(ctx)
	_logger, formatCaller, vCaller := addCaller(withContext(ctx))
	addDebugMessage(ctx, fmt.Sprintf("[PANIC] "+formatCaller, vCaller...))
	addDebugMessage(ctx, fmt.Sprintf(format, v...))
	_v := append(vField, v...)
	v = append(vCaller, _v...)
	if prefix, ok := ctx.Value(loggerPrefixKey).(string); ok {
		format = prefix + "\t" + format
	}
	_logger.Panic(fmt.Sprintf(formatCaller+formatField+newLine+format, v...))
}

func FatalfCtx(ctx context.Context, format string, v ...interface{}) {
	formatField, vField := addExField(ctx)
	_logger, formatCaller, vCaller := addCaller(withContext(ctx))
	addDebugMessage(ctx, fmt.Sprintf("[FATAL] "+formatCaller, vCaller...))
	addDebugMessage(ctx, fmt.Sprintf(format, v...))
	_v := append(vField, v...)
	v = append(vCaller, _v...)
	if prefix, ok := ctx.Value(loggerPrefixKey).(string); ok {
		format = prefix + "\t" + format
	}
	_logger.Fatal(fmt.Sprintf(formatCaller+formatField+newLine+format, v...))
}
