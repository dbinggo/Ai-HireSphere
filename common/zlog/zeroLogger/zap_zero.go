package zeroLogger

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type zeroLogger struct {
	logger     *zap.Logger
	formatJson bool
	newLine    string
}

var _ logx.Writer = (*zeroLogger)(nil) // 接口实现检查
func NewZeroLogger(logger *zap.Logger, formatJson bool, newLine string) logx.Writer {
	return &zeroLogger{logger: logger, formatJson: formatJson, newLine: newLine}
}

func (l *zeroLogger) formatSeparator() string {
	if l.formatJson {
		return "%v"
	} else {
		return "%v\t"
	}

}

func (l *zeroLogger) Alert(v interface{}) {
	l.logger.Error(fmt.Sprint(v))
}

func (l *zeroLogger) Close() error {
	return l.logger.Sync()
}

func (l *zeroLogger) Debug(v interface{}, fields ...logx.LogField) {
	if l.formatJson {
		l.logger.Debug(fmt.Sprint(v), toZapFields(fields...)...)
	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(l.formatSeparator(), f.Value)
		}
		l.logger.Debug(ex + l.newLine + fmt.Sprint(v))
	}
}

func (l *zeroLogger) Error(v interface{}, fields ...logx.LogField) {
	if l.formatJson {
		l.logger.Error(fmt.Sprint(v), toZapFields(fields...)...)
	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(l.formatSeparator(), f.Value)
		}
		l.logger.Error(ex + l.newLine + fmt.Sprint(v))
	}
}

func (l *zeroLogger) Info(v interface{}, fields ...logx.LogField) {
	if l.formatJson {
		l.logger.Info(fmt.Sprint(v), toZapFields(fields...)...)
	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(l.formatSeparator(), f.Value)
		}
		l.logger.Info(ex + l.newLine + fmt.Sprint(v))
	}
}

func (l *zeroLogger) Severe(v interface{}) {
	l.logger.Error(fmt.Sprint(v))
}

func (l *zeroLogger) Slow(v interface{}, fields ...logx.LogField) {
	if l.formatJson {
		l.logger.Warn(fmt.Sprint(v), toZapFields(fields...)...)
	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(l.formatSeparator(), f.Value)
		}
		l.logger.Warn(ex + l.newLine + fmt.Sprint(v))
	}
}

func (l *zeroLogger) Stack(v interface{}) {
	l.logger.Error(fmt.Sprint(v), zap.Stack("stack"))
}

func (l *zeroLogger) Stat(v interface{}, fields ...logx.LogField) {
	if l.formatJson {
		l.logger.Info(fmt.Sprint(v), toZapFields(fields...)...)

	} else {
		ex := ""
		for _, f := range fields {
			ex += fmt.Sprintf(l.formatSeparator(), f.Value)
		}
		l.logger.Info(ex + l.newLine + fmt.Sprint(v))

	}
}

func toZapFields(fields ...logx.LogField) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zapFields = append(zapFields, zap.Any(f.Key, f.Value))
	}
	return zapFields
}
