package idataaccess

import (
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/common/model/enums"
	"context"
	"github.com/dbinggo/gerr"
)

// 本类主要是关于User的数据库操作
type IUserGorm interface {
	SaveUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, gerr.Error)
	FindUserById(ctx context.Context, id int64) (*entity.UserEntity, gerr.Error)
	FindUserByLoginType(ctx context.Context, loginType enums.UserRegisterWayType, data string) (*entity.UserEntity, gerr.Error)
}
