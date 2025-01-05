package idataaccess

import (
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"context"
)

// 本类主要是关于Resume的数据库操作
type IUserGorm interface {
	SaveUser(ctx context.Context, user *entity.User) (*entity.User, error)
	FindUserById(ctx context.Context, id int64) (*entity.User, error)
}
