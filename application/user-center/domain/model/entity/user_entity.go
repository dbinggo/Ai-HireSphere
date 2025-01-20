package entity

import (
	"Ai-HireSphere/common/model/enums"
	"Ai-HireSphere/common/utils"
	"Ai-HireSphere/common/utils/jwt"
	"errors"
	"github.com/dbinggo/gerr"
)

// 本层不允许引用repo层 否则会造成循环依赖

// 用户基本信息
type User struct {
	Id       int64
	Sex      enums.UserSex
	Role     enums.UserRole
	Avatar   string
	UserName string
	Email    string
	Phone    string
}

// 充血模型

func (u *User) Register(way enums.UserRegisterWayType) gerr.Error {
	if err := u.CheckRegister(way); err != nil {
		return err
	}
	return nil
}

// Validate
//
//	@Description: 校验用户注册信息
//	@receiver u
//	@return error
func (u *User) Validate() error {
	var (
		errEmail error
		errPhone error
	)
	if u.Email != "" {
		errEmail = utils.CheckEmail(u.Email)
	}
	if u.Phone != "" {
		errPhone = utils.CheckPhone(u.Phone)
	}
	return errors.Join(errEmail, errPhone)
}

// CheckRegister
//
//	@Description: 用户注册前置校验 目前只有登陆方式校验
//	@receiver u
//	@param way
//	@return error
func (u *User) CheckRegister(way enums.UserRegisterWayType) gerr.Error {
	switch way {
	case enums.UserRegisterWayTypeEmail:
		// 校验email邮箱是否正确
		err := utils.CheckEmail(u.Email)
		if err != nil {
			return err
		}
	case enums.UserRegisterWayTypePhone:
		// 校验手机号是否正确
		err := utils.CheckPhone(u.Phone)
		if err != nil {
			return err
		}
	}
	return nil
}

// GenerateToken
//
//	@Description: 生成用户Token
//	@receiver u
//	@return string
//	@return error
func (u *User) GenerateToken() (string, gerr.Error) {
	token, err := jwt.GenerateToken(u.Id)
	if err != nil {
		err = gerr.WrapSysErrf(err, "生成token失败")
		return "", err.(gerr.Error)
	}
	return token, nil
}
