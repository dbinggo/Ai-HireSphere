package repository

import (
	"Ai-HireSphere/application/user-center/domain/irepository"
	dataaccess "Ai-HireSphere/application/user-center/infrastructure/repository/data_access"
	redisaccess "Ai-HireSphere/application/user-center/infrastructure/repository/redis_access"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"sync"
)

// Repo is a global variable that is an instance of RepoBroker
var (
	repo irepository.IRepoBroker
	once sync.Once
)

type RepoStruct struct {
	*dataaccess.GormOpts
	*redisaccess.RedisOpts
}

// 初始化仓储服务
func NewRepoBroker(db *gorm.DB, cli *redis.Client) irepository.IRepoBroker {
	// 单例模式 初始化仓储服务
	once.Do(func() {
		repo = &RepoStruct{
			GormOpts:  dataaccess.NewGormOpts(db),
			RedisOpts: redisaccess.NewRedisOpts(cli),
		}
	})
	return repo
}
