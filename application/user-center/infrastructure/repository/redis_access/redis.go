package redisaccess

import (
	ireidsaccess "Ai-HireSphere/application/user-center/domain/irepository/ireids_access"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

// RedisOpts ...
type RedisOpts struct {
	cli *redis.Client
}

var _ ireidsaccess.IRedisAccess = (*RedisOpts)(nil)

func NewRedisOpts(cli *redis.Client) *RedisOpts {
	return &RedisOpts{
		cli: cli,
	}
}

func (o *RedisOpts) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	return o.cli.Set(ctx, key, value, expiration).Err()
}

func (o *RedisOpts) Get(ctx context.Context, key string) (interface{}, error) {
	return o.cli.Get(ctx, key).Result()
}

// HSet ... notice that value should be marshalizable.
func (o *RedisOpts) HSet(ctx context.Context, key string, value any, expiration time.Duration) error {
	var (
		b   []byte
		err error
	)

	// 序列化 value
	if b, err = json.Marshal(value); err != nil {

	}
	// 存入
	if _, err = o.cli.Set(ctx, key, b, expiration).Result(); err != nil {
	}

	// RET
	return nil
}

// HGet ... notice that receiver should be a pointer to the designated instance (mostly map or struct) that will store the fetched value.
func (o *RedisOpts) HGet(ctx context.Context, key string, receiver any) error {
	var (
		str string
		err error
	)

	// 取出
	if str, err = o.cli.Get(ctx, key).Result(); err != nil {
		return err
	}
	// 反序列化
	if err = json.Unmarshal([]byte(str), receiver); err != nil {
		return err
	}

	// RET
	return nil
}

// Del ...
func (o *RedisOpts) Del(ctx context.Context, key string) error {
	_, err := o.cli.Del(ctx, key).Result()
	if err != nil {
		return err
	}

	// RET
	return nil
}

// ExistsKeyRedis ...
func (o *RedisOpts) ExistsKeyRedis(ctx context.Context, key string) (bool, error) {
	exists, err := o.cli.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}

	if exists == 1 {
		return true, nil
	}

	// RET
	return false, nil
}

// SetKeyTTLInRedis ...
func (o *RedisOpts) SetKeyTTLInRedis(ctx context.Context, key string, expiration time.Duration) (bool, error) {
	ok, err := o.cli.Expire(ctx, key, expiration).Result()
	if err != nil {
		return false, err
	}

	// RET
	return ok, nil
}

// Lock ...
func (o *RedisOpts) Lock(ctx context.Context, key string, timeout time.Duration) (bool, error) {
	var luaScript = redis.NewScript(`
		if redis.call("SETNX", KEYS[1], ARGV[1]) == 1 then
			return redis.call("PEXPIRE", KEYS[1], ARGV[2])
		end
		return 0
	`)
	success, err := luaScript.Run(ctx, o.cli, []string{key}, 1, int64(timeout/time.Millisecond)).Result()
	if err != nil {
		return false, err
	}

	lockAcquired := success.(int64) == 1
	return lockAcquired, nil
}

// Unlock ...
func (o *RedisOpts) Unlock(ctx context.Context, key string) (bool, error) {
	luaScript := redis.NewScript(`
		if redis.call("GET", KEYS[1]) == ARGV[1] then
			return redis.call("DEL", KEYS[1])
		else
			return 0
		end
	`)
	_, err := luaScript.Run(ctx, o.cli, []string{key}, 1).Result()
	if err != nil {
		return false, err
	}
	return true, nil
}
