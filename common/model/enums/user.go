package enums

// 类型枚举值，禁止使用字符串进行类型定义
// 命名要求：类名+字段名+Type+实际含义

// 避免不同类型在不同服务中定义不同，此枚举特地放在公共包中 规范类枚举
type UserSex uint

const (
	UserSexTypeUnknown UserSex = 0
	UserSexTypeMale    UserSex = 1
	UserSexTypeFemale  UserSex = 2
)

type UserRole string

const (
	UserRoleTypeAdmin UserRole = "admin"
	UserRoleTypeUser  UserRole = "user"

	UserRoleTypeDefault UserRole = UserRoleTypeUser
)

type UserRegisterWayType string

const (
	UserRegisterWayTypeEmail UserRegisterWayType = "email"
	UserRegisterWayTypePhone UserRegisterWayType = "phone"
)
