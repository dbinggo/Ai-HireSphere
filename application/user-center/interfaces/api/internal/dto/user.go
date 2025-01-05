package dto

import (
	"Ai-HireSphere/application/user-center/domain/model/entity"
	"Ai-HireSphere/application/user-center/interfaces/api/internal/types"
	"Ai-HireSphere/common/model/enums"
)

func RegisterReq(req *types.RegisterReq) (way enums.UserRegisterWayType, user *entity.User) {
	way = enums.UserRegisterWayType(req.Way)

	user = &entity.User{
		Role: enums.UserRoleTypeDefault,
	}

	switch way {
	case enums.UserRegisterWayTypeEmail:
		user.Email = req.Data
	case enums.UserRegisterWayTypePhone:
		user.Phone = req.Data
	}
	return way, user
}
