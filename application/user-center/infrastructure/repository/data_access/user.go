package dataaccess

import (
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/common/model"
	"context"
)

// IResumeGorm is an interface that defines the methods for accessing data from MySQL
var _ idataaccess.IUserGorm = (*GormOpts)(nil)

// 所有方法返回值不应该返回数据库模型，
//即：底层数据库存储对上层领域模型没有任何感知，你想怎么存就怎么存，
//弱化数据库概念，数据库只是适配器

type userGorm struct {
}

func (o *GormOpts) SaveUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	u := o.userEntryToModel(user)
	u, err := u.Create(o.db, u)
	if err != nil {
		return nil, err
	}
	return o.userModelToEntry(u), nil
}

func (o *GormOpts) FindUserById(ctx context.Context, id int64) (*entity.User, error) {
	z, err := (&model.TUser{}).GetOne(o.db, "id = ?", id)
	if err != nil {
		return nil, err
	}
	return o.userModelToEntry(z), nil
}

// todo 待完善
func (o *GormOpts) userModelToEntry(user model.TUser) *entity.User {
	return &entity.User{}
}
func (o *GormOpts) userEntryToModel(user *entity.User) model.TUser {
	return model.TUser{}
}
