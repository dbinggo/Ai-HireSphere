package decorator

import (
	"context"

	"github.com/sirupsen/logrus"
)

// ApplyQueryDecorators is a function that applies decorators to a query handler.
// H:QueryHandler的query参数类型
// R:QueryHandler的返回值类型
func ApplyQueryDecorators[H any, R any](handler QueryHandler[H, R], logger *logrus.Entry, metricsClient MetricsClient) QueryHandler[H, R] {
	return queryLoggingDecorator[H, R]{
		base: queryMetricsDecorator[H, R]{
			base:   handler,
			client: metricsClient,
		},
		logger: logger,
	}
}

type QueryHandler[Q any, R any] interface {
	Handle(ctx context.Context, q Q) (R, error)
}
