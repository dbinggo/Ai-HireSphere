package ireidsaccess

import (
	"context"
	"time"
)

// IRedisAccess defines the set of methods that we use for accessing data from Redis
type IRedisAccess interface {
	// 基础方法
	Set(ctx context.Context, key string, value any, expiration time.Duration) error
	Get(ctx context.Context, key string, receiver any) error
	ExistsKeyRedis(ctx context.Context, key string) (bool, error)
	SetKeyTTLInRedis(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Del(ctx context.Context, key string) error
	Lock(ctx context.Context, key string, expiration time.Duration) (bool, error)
	Unlock(ctx context.Context, key string) (bool, error)
}
