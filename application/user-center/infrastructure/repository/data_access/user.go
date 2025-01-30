package dataaccess

import (
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/common/codex"
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

func (o *GormOpts) SaveUser(ctx context.Context, user *entity.UserEntity) (*entity.UserEntity, gerr.Error) {

	userModel := user.Transform()

	userModel, err := userModel.Save(ctx, o.db, userModel)
	if err != nil {
		err = gerr.Wraps(codex.ServerErr, err)
		return nil, err.(gerr.Error)
	}
	return user.From(userModel), nil
}

func (o *GormOpts) FindUserById(ctx context.Context, id int64) (user *entity.UserEntity, err2 gerr.Error) {
	z, err := (&model.CommonAdapter[model.TUser]{}).GetOne(ctx, o.db, "id = ?", id)
	if err != nil {
		err = gerr.Wraps(codex.ServerErr, err)
		return nil, err.(gerr.Error)
	}
	return (&entity.UserEntity{}).From(z), nil
}

func (o *GormOpts) FindUserByLoginType(ctx context.Context, loginType enums.UserRegisterMethodType, data string) (user *entity.UserEntity, err gerr.Error) {
	if loginType == enums.UserRegisterWayTypeEmail {
		z, err := (&model.CommonAdapter[model.TUser]{}).GetOne(ctx, o.db, "email = ?", data)
		if err != nil {
			err = gerr.Wraps(codex.ServerErr, err)
			return nil, err.(gerr.Error)
		}
		user.From(z)
		return user, nil
	} else {
		z, err := (&model.CommonAdapter[model.TUser]{}).GetOne(ctx, o.db, "phone = ?", data)
		if err != nil {
			err = gerr.Wraps(codex.ServerErr, err)
			return nil, err.(gerr.Error)
		}
		return (&entity.UserEntity{}).From(z), nil
	}
}
