package dbLogger

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	gormLogger "gorm.io/gorm/logger"
	"time"
)

// 适配gorm 的gormLogger接口
//type Interface interface {
//	LogMode(LogLevel) Interface
//	Info(context.Context, string, ...interface{})
//	Warn(context.Context, string, ...interface{})
//	Error(context.Context, string, ...interface{})
//	Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error)
//}

type dbLog struct {
	LogLevel gormLogger.LogLevel
}

var _ gormLogger.Interface = (*dbLog)(nil) //接口实现检查
func New() *dbLog {
	return new(dbLog)
}

func (l *dbLog) LogMode(level gormLogger.LogLevel) gormLogger.Interface {
	l.LogLevel = level
	return l
}

func (l *dbLog) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < gormLogger.Info {
		return
	}
	logx.WithContext(ctx).Debugf(msg, data)
}
func (l *dbLog) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < gormLogger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(msg, data)
}

func (l *dbLog) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel < gormLogger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(msg, data)
}

func (l *dbLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	//这块的逻辑可以自己根据业务情况修改
	//fmt.Println(l.LogLevel)
	elapsed := time.Since(begin)
	sql, rows := fc()
	logx.WithContext(ctx).WithDuration(elapsed).Slowf("Trace sql: %v  row： %v  err: %v", sql, rows, err)
}
