package repositor

import (
	"Ai-HireSphere/application/interview/domain/irepository"
	dataaccess "Ai-HireSphere/application/interview/infrastructure/repository/data_access"
	"Ai-HireSphere/common/coze"
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
	*coze.CozeApi
}

// 初始化仓储服务
func NewRepoBroker(db *gorm.DB, cli *redis.Client) irepository.IRepoBroker {
	// 单例模式 初始化仓储服务
	once.Do(func() {
		repo = &RepoStruct{
			GormOpts: dataaccess.NewGormOpts(db),
		}
	})
	return repo
}
