package enums

// 类型枚举值，禁止使用字符串进行类型定义
// 命名要求：类名+字段名+Type+实际含义

type UserSex uint

const (
	UserSexTypeMale   UserSex = 1
	UserSexTypeFemale UserSex = 2
)
