package model

import (
	"Ai-HireSphere/common/model/enums"
	"gorm.io/gorm"
)

type Example struct {
	CommonModel
	Name string        //	姓名
	Sex  enums.UserSex // 1: 男 2: 女 enums.UserSex
	Age  int           // 年龄
}

// 类型断言
var _ ICommonModel = &Example{}

func NewExample() *Example {
	return &Example{}
}

// 使用的是充血模型
// 可以扩展更多方法，本类只是一个示例
func (e Example) TableName() string {
	return "t_example"
}

func (e *Example) GetOne(db *gorm.DB, where interface{}) (Example, error) {
	return _getOne[Example](db, where)
}

func (e *Example) GetMulti(db *gorm.DB, where interface{}) ([]Example, error) {
	return _getMulti[Example](db, e.TableName(), where)
}

func (e *Example) Create(db *gorm.DB, data Example) error {
	return _create[Example](db, data)
}

func (e *Example) UpdateOne(db *gorm.DB, data Example, selects ...string) (Example, error) {
	return _updateOne[Example](db, data, selects...)
}

func (e *Example) Delete(db *gorm.DB, ids ...[]int) error {
	return _delete[Example](db, ids...)
}
