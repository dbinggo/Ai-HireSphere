package repository

import (
	"Ai-HireSphere/common/repository/data_access"
	"Ai-HireSphere/common/repository/redis_access"
)

// RepoBroker is a should-be-static, global-accessible RepoBroker, used in service.NewService()
type RepoBroker interface {
	dataaccess.MysqlInterface
	redisaccess.RedisInterface
}

var Repo RepoBroker = NewRepoBroker()

// RepoStruct is a trivial implementation of RepoBroker, embedding every single struct under package repository.
type RepoStruct struct {
	*dataaccess.MysqlOpts
	*redisaccess.RedisOpts
}

// NewRepoBroker returns a RepoBroker
func NewRepoBroker() RepoBroker {
	return &RepoStruct{
		MysqlOpts: new(dataaccess.MysqlOpts),
		RedisOpts: new(redisaccess.RedisOpts),
	}
}
