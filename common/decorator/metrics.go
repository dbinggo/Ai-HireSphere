package decorator

import (
	"context"
	"fmt"
	"strings"
	"time"
)

// 装饰器模式 添加监控 可以先直接打印，后续对接监控
type MetricsClient interface {
	Inc(key string, value int)
}

type commandMetricsDecorator[C any] struct {
	base   CommandHandler[C]
	client MetricsClient
}

// 使用装饰器模式 添加命令的监控
func (d commandMetricsDecorator[C]) Handle(ctx context.Context, cmd C) (err error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(cmd))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("commands.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("commands.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("commands.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, cmd)
}

type queryMetricsDecorator[C any, R any] struct {
	base   QueryHandler[C, R]
	client MetricsClient
}

// 使用装饰器模式 添加查询的监控
func (d queryMetricsDecorator[C, R]) Handle(ctx context.Context, query C) (result R, err error) {
	start := time.Now()

	actionName := strings.ToLower(generateActionName(query))

	defer func() {
		end := time.Since(start)

		d.client.Inc(fmt.Sprintf("querys.%s.duration", actionName), int(end.Seconds()))

		if err == nil {
			d.client.Inc(fmt.Sprintf("querys.%s.success", actionName), 1)
		} else {
			d.client.Inc(fmt.Sprintf("querys.%s.failure", actionName), 1)
		}
	}()

	return d.base.Handle(ctx, query)
}
