package dataaccess

import (
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/common/model"
	"Ai-HireSphere/common/model/enums"
	"context"
	"github.com/dbinggo/gerr"
)

// IResumeGorm is an interface that defines the methods for accessing data from MySQL
var _ idataaccess.IUserGorm = (*GormOpts)(nil)

// 所有方法返回值不应该返回数据库模型，
//即：底层数据库存储对上层领域模型没有任何感知，你想怎么存就怎么存，
//弱化数据库概念，数据库只是适配器

type userGorm struct {
}

func (o *GormOpts) SaveUser(ctx context.Context, user *entity.User) (*entity.User, gerr.Error) {
	u := o.userEntryToModel(user)
	u, err := u.Create(o.db, u)
	if err != nil {
		err = gerr.DefaultSysErr()
		return nil, err.(gerr.Error)
	}
	return o.userModelToEntry(u), nil
}

func (o *GormOpts) FindUserById(ctx context.Context, id int64) (*entity.User, gerr.Error) {
	z, err := (&model.TUser{}).GetOne(o.db, "id = ?", id)
	if err != nil {
		err = gerr.DefaultSysErr()
		return nil, err.(gerr.Error)
	}
	return o.userModelToEntry(z), nil
}

func (o *GormOpts) FindUserByLoginType(ctx context.Context, loginType enums.UserRegisterWayType, data string) (*entity.User, gerr.Error) {
	if loginType == enums.UserRegisterWayTypeEmail {
		z, err := (&model.TUser{}).GetOne(o.db, "email = ?", data)
		if err != nil {
			err = gerr.DefaultSysErr()
			return nil, err.(gerr.Error)
		}
		return o.userModelToEntry(z), nil
	} else {
		z, err := (&model.TUser{}).GetOne(o.db, "phone = ?", data)
		if err != nil {
			err = gerr.DefaultSysErr()
			return nil, err.(gerr.Error)
		}
		return o.userModelToEntry(z), nil
	}
	return nil, gerr.DefaultSysErr()
}

func (o *GormOpts) userModelToEntry(user model.TUser) *entity.User {
	return &entity.User{}
}
func (o *GormOpts) userEntryToModel(user *entity.User) model.TUser {
	return model.TUser{}
}
