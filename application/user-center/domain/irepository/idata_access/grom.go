package idataaccess

// IDataAccess defines the set of methods that we use for accessing data from MySQL

type IDataAccess interface {
	IUserGorm
	//ITransaction[T]
}

// 实现基础事务功能
type ITransaction[T any] interface {
	Transaction(fc func(tx *T) error)
}
