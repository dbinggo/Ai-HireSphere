package entity

type ICommonEntity[T any, F any] interface {
	// Transform 转换为数据库schema结构
	Transform() F
	// From 从数据库schema结构转换为实体
	From(F) T
}
