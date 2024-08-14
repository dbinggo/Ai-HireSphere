package zlog

import (
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"runtime"
	"slices"
)

const (
	logger_formate_json = "json"

	loggerFieldKey  = "field"
	loggerDebugKey  = "debug"
	loggerCallerKey = "caller"
	loggerSpanKey   = "span"
	loggerTraceKey  = "trace"

	loggerPrefixKey = "prefix"
	// 分隔符

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
var zlogConfig *ZlogConfig = &ZlogConfig{}
var newLine = "\n"

func SetZlog(config ZlogConfig) {
	zlogConfig = &config
	if !config.NewLine {
		newLine = ""
	}
}
func SetPrefix(ctx *context.Context, prefix string) context.Context {
	if *ctx == nil {
		*ctx = context.Background()
	}
	*ctx = context.WithValue(*ctx, loggerPrefixKey, prefix)
	return *ctx
}
func getLogger() zap.Logger {
	return *logger
}

type colour int

const (
	ColourBlack colour = iota
	ColourRed
	ColourGreen
	ColourYellow
	ColourBlue
	ColourPurple
	ColourCyan
	ColourWhite
)

// SetColour 设置字体颜色 30黑 31红 32绿 33黄 34蓝 35紫 36青 37白
func SetColour(text string, colour colour) string {
	if !needColour() {
		return text
	}
	return fmt.Sprintf("\u001b[;%dm%s\u001b[m", 30+colour, text)

}

// SetBlackColour 设置背景颜色 40黑 41红 42绿 43黄 44蓝 45紫 46青 47白
func SetBlackColour(text string, colour colour) string {
	if !needColour() {
		return text
	}
	return fmt.Sprintf("\u001b[;%dm%s\u001b[0m", 40+colour, text)

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

	if fieldOld, ok := ctx.Value(loggerFieldKey).([]zapcore.Field); ok {
		var _newField map[string]zapcore.Field
		var __newField []zapcore.Field
		for _, field := range fields {
			_newField[field.Key] = field
			__newField = append(__newField, field)
		}

		// 如果存在该key 只做更新操作 不做添加操作
		var fieldNew []zapcore.Field
		for _, f := range fieldOld {
			if !slices.Contains(__newField, f) {
				fieldNew = append(fieldNew, f)
			} else {
				fieldNew = append(fieldNew, _newField[f.Key])
			}
		}
		ctx = context.WithValue(ctx, loggerFieldKey, fieldNew)
	} else {
		ctx = context.WithValue(ctx, loggerFieldKey, fields)
	}

	if debug() {
		if _, ok := ctx.Value(loggerDebugKey).([]string); !ok {
			ctx = context.WithValue(ctx, loggerDebugKey, []string{})
		}

	}
	return ctx
}
func initLogger(zapLogger *zap.Logger) {
	logger = zapLogger
}

func Infof(format string, v ...interface{}) {
	_logger := getLogger()
	_logger, textCaller := addCaller(&_logger)
	_logger.Info(fmt.Sprintf(textCaller+newLine+format, v...))
}

func Errorf(format string, v ...interface{}) {
	_logger := getLogger()
	_logger, textCaller := addCaller(&_logger)
	_logger.Error(fmt.Sprintf(textCaller+newLine+format, v...))
}

func Warnf(format string, v ...interface{}) {
	_logger := getLogger()
	_logger, textCaller := addCaller(&_logger)
	_logger.Warn(fmt.Sprintf(textCaller+newLine+format, v...))
}

func Debugf(format string, v ...interface{}) {
	_logger := getLogger()
	_logger, textCaller := addCaller(&_logger)
	_logger.Debug(fmt.Sprintf(textCaller+newLine+format, v...))
}

func Panicf(format string, v ...interface{}) {
	_logger := getLogger()
	_logger, textCaller := addCaller(&_logger)
	_logger.Panic(fmt.Sprintf(textCaller+newLine+format, v...))
}

func Fatalf(format string, v ...interface{}) {
	_logger := getLogger()
	_logger, textCaller := addCaller(&_logger)
	_logger.Fatal(fmt.Sprintf(textCaller+newLine+format, v...))
}

func addCaller(_logger *zap.Logger) (zap.Logger, string) {
	format := "%s:%d"
	_, file, line, _ := runtime.Caller(zlogConfig.CallerSkip)
	_v := make([]interface{}, 0)
	file = file[len(zlogConfig.Path)+1:]
	_v = append(_v, file, line)
	if formatJson() {
		_logger = _logger.With(zap.String(loggerCallerKey, fmt.Sprintf(format, file, line)))
		return *_logger, ""
	}
	return *logger, fmt.Sprintf(format+getT(), _v...)
}
func addSpan(ctx context.Context, _logger *zap.Logger) (zap.Logger, string) {

	spanId := trace.SpanIDFromContext(ctx)
	if spanId == "" {
		return *_logger, ""
	}
	if formatJson() {
		_logger = _logger.With(zap.String(loggerSpanKey, spanId))
		return *_logger, ""
	}
	format := "%v" + getT()
	return *logger, fmt.Sprintf(format, spanId)

}
func addTrace(ctx context.Context, _logger *zap.Logger) (zap.Logger, string) {
	traceId := trace.TraceIDFromContext(ctx)
	if traceId == "" {
		return *_logger, ""
	}
	if formatJson() {
		_logger = _logger.With(zap.String(loggerTraceKey, traceId))
		return *_logger, ""
	}
	format := "%v" + getT()
	return *logger, fmt.Sprintf(format, traceId)
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
func getT() string {
	if formatJson() {
		return ""
	} else {
		return "\t"
	}
}
func addExField(ctx context.Context, _logger *zap.Logger) (zap.Logger, string) {

	if exField, ok := ctx.Value(loggerFieldKey).([]zapcore.Field); ok {
		if formatJson() {
			_logger = _logger.With(exField...)
			return *_logger, ""
		} else {
			format := "%v" + getT()
			ret := ""
			for _, field := range exField {
				ret += fmt.Sprintf(format, field.String)
			}
			return *_logger, ret
		}
	}
	return *_logger, ""

}
func addPrefix(ctx context.Context, _logger *zap.Logger) (zap.Logger, string) {
	if prefix, ok := ctx.Value(loggerPrefixKey).(string); ok {
		if formatJson() {
			_logger = _logger.With(zap.String(loggerPrefixKey, prefix))
			return *_logger, ""
		}
		return *logger, prefix + getT()
	}
	return *logger, ""

}

// 下面的logger方法会携带trace id

func InfofCtx(ctx context.Context, format string, v ...interface{}) {
	_logger := getLogger()
	_logger, caller := addCaller(&_logger)
	_logger, traceId := addTrace(ctx, &_logger)
	_logger, spanId := addSpan(ctx, &_logger)
	_logger, field := addExField(ctx, &_logger)
	_logger, prefix := addPrefix(ctx, &_logger)
	addDebugMessage(ctx, fmt.Sprintf("[INFO]  "+caller+traceId+spanId+field))
	addDebugMessage(ctx, fmt.Sprintf(prefix+format, v...))
	_logger.Info(fmt.Sprintf(caller+traceId+spanId+field+newLine+prefix+format, v...))
}

func ErrorfCtx(ctx context.Context, format string, v ...interface{}) {
	_logger := getLogger()
	_logger, caller := addCaller(&_logger)
	_logger, traceId := addTrace(ctx, &_logger)
	_logger, spanId := addSpan(ctx, &_logger)
	_logger, field := addExField(ctx, &_logger)
	_logger, prefix := addPrefix(ctx, &_logger)
	addDebugMessage(ctx, fmt.Sprintf("[ERROR] "+caller+traceId+spanId+field))
	addDebugMessage(ctx, fmt.Sprintf(prefix+format, v...))
	_logger.Error(fmt.Sprintf(caller+traceId+spanId+field+newLine+prefix+format, v...))
}

func WarnfCtx(ctx context.Context, format string, v ...interface{}) {
	_logger := getLogger()
	_logger, caller := addCaller(&_logger)
	_logger, traceId := addTrace(ctx, &_logger)
	_logger, spanId := addSpan(ctx, &_logger)
	_logger, field := addExField(ctx, &_logger)
	_logger, prefix := addPrefix(ctx, &_logger)
	addDebugMessage(ctx, fmt.Sprintf("[WARN]  "+caller+traceId+spanId+field))
	addDebugMessage(ctx, fmt.Sprintf(prefix+format, v...))
	_logger.Warn(fmt.Sprintf(caller+traceId+spanId+field+newLine+prefix+format, v...))
}

func DebugfCtx(ctx context.Context, format string, v ...interface{}) {
	_logger := getLogger()
	_logger, caller := addCaller(&_logger)
	_logger, traceId := addTrace(ctx, &_logger)
	_logger, spanId := addSpan(ctx, &_logger)
	_logger, field := addExField(ctx, &_logger)
	_logger, prefix := addPrefix(ctx, &_logger)
	addDebugMessage(ctx, fmt.Sprintf("[DEBUG] "+caller+traceId+spanId+field))
	addDebugMessage(ctx, fmt.Sprintf(prefix+format, v...))
	_logger.Debug(fmt.Sprintf(caller+traceId+spanId+field+newLine+prefix+format, v...))
}

func PanicfCtx(ctx context.Context, format string, v ...interface{}) {
	_logger := getLogger()
	_logger, caller := addCaller(&_logger)
	_logger, traceId := addTrace(ctx, &_logger)
	_logger, spanId := addSpan(ctx, &_logger)
	_logger, field := addExField(ctx, &_logger)
	_logger, prefix := addPrefix(ctx, &_logger)
	addDebugMessage(ctx, fmt.Sprintf("[PANIC] "+caller+traceId+spanId+field))
	addDebugMessage(ctx, fmt.Sprintf(prefix+format, v...))
	_logger.Panic(fmt.Sprintf(caller+traceId+spanId+field+newLine+prefix+format, v...))
}

func FatalfCtx(ctx context.Context, format string, v ...interface{}) {
	_logger := getLogger()
	_logger, caller := addCaller(&_logger)
	_logger, traceId := addTrace(ctx, &_logger)
	_logger, spanId := addSpan(ctx, &_logger)
	_logger, field := addExField(ctx, &_logger)
	_logger, prefix := addPrefix(ctx, &_logger)
	addDebugMessage(ctx, fmt.Sprintf("[FATAL] "+caller+traceId+spanId+field))
	addDebugMessage(ctx, fmt.Sprintf(prefix+format, v...))
	_logger.Fatal(fmt.Sprintf(caller+traceId+spanId+field+newLine+prefix+format, v...))
}

// todo 实现logx的接口
