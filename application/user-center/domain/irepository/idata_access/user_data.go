package idataaccess

import (
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/common/model/enums"
	"context"
	"github.com/dbinggo/gerr"
)

// 本类主要是关于Resume的数据库操作
type IUserGorm interface {
	SaveUser(ctx context.Context, user *entity.User) (*entity.User, gerr.Error)
	FindUserById(ctx context.Context, id int64) (*entity.User, gerr.Error)
	FindUserByLoginType(ctx context.Context, loginType enums.UserRegisterWayType, data string) (*entity.User, gerr.Error)
}
