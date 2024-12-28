package dataaccess

// MysqlInterface 所有数据库方法的统一接口
type MysqlInterface interface {
	getExample
}

// MysqlOpts 实现了所有数据库方法的统一接口
type MysqlOpts struct{}
