package zlog

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

const callerSkipOffset = 3

var formatSeparator = "%v" + getT()

type ZapWriter struct {
	logger *zap.Logger
}

func NewZapWriter(logger *zap.Logger) logx.Writer {
	return &ZapWriter{logger: logger}
}

func (w *ZapWriter) Alert(v interface{}) {
	w.logger.Error(fmt.Sprint(v))
}

func (w *ZapWriter) Close() error {
	return w.logger.Sync()
}

func (w *ZapWriter) Debug(v interface{}, fields ...logx.LogField) {
	if formatJson() {
		w.logger.Debug(fmt.Sprint(v), toZapFields(fields...)...)
	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(formatSeparator, f.Value)
		}
		w.logger.Debug(ex + newLine + fmt.Sprint(v))
	}
}

func (w *ZapWriter) Error(v interface{}, fields ...logx.LogField) {
	if formatJson() {
		w.logger.Error(fmt.Sprint(v), toZapFields(fields...)...)
	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(formatSeparator, f.Value)
		}
		w.logger.Error(ex + newLine + fmt.Sprint(v))
	}
}

func (w *ZapWriter) Info(v interface{}, fields ...logx.LogField) {
	if formatJson() {
		w.logger.Info(fmt.Sprint(v), toZapFields(fields...)...)
	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(formatSeparator, f.Value)
		}
		w.logger.Info(ex + newLine + fmt.Sprint(v))
	}
}

func (w *ZapWriter) Severe(v interface{}) {
	w.logger.Error(fmt.Sprint(v))
}

func (w *ZapWriter) Slow(v interface{}, fields ...logx.LogField) {
	if formatJson() {
		w.logger.Warn(fmt.Sprint(v), toZapFields(fields...)...)
	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(formatSeparator, f.Value)
		}
		w.logger.Warn(ex + newLine + fmt.Sprint(v))
	}
}

func (w *ZapWriter) Stack(v interface{}) {
	w.logger.Error(fmt.Sprint(v), zap.Stack("stack"))
}

func (w *ZapWriter) Stat(v interface{}, fields ...logx.LogField) {
	if formatJson() {
		w.logger.Info(fmt.Sprint(v), toZapFields(fields...)...)

	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(formatSeparator, f.Value)
		}
		w.logger.Info(ex + newLine + fmt.Sprint(v))

	}
}

func toZapFields(fields ...logx.LogField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
