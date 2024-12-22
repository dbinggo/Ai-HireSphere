package dataaccess

import (
	"Ai-HireSphere/common/model"
	"gorm.io/gorm"
)

type getExample interface {
	Calculate(ids []int16) int
}

// Calculate 对传进来的example id进行计算 返回计算结果
func (m *MysqlOpts) Calculate(ids []int16) int {
	// 调用方法对象 无其他作用
	t := model.NewExample()
	db := &gorm.DB{}
	findMap := make(map[string]interface{})
	findMap["id"] = ids
	ret, err := t.GetMulti(db, findMap)
	if err != nil {
		return 0
	}
	ans := 0
	for _, v := range ret {
		ans += v.Age
	}
	return int(ans)

}
