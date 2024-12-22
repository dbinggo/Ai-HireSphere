package redisaccess

type getExampleInterface interface {
	GetAge(id int) int
}

// Calculate 对传进来的example id进行计算 返回计算结果
func (m *RedisOpts) GetAge(id int) int {
	// 调用方法对象 无其他作用
	return 0
}
