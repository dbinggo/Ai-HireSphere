package model

import (
	"gorm.io/gorm"
	"time"
)

// 本类文件提供给最基础的增删改查语句，便于扩展
type CommonModel struct {
	ID        int64 `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type ICommonModel interface {
	TableName() string
}

// 数据库接口适配器，提供最基本CURD能力
type IDBAdapter[T ICommonModel] interface {
	GetOne(db *gorm.DB, where ...interface{}) (*T, error)
	GetMulti(db *gorm.DB, where ...interface{}) (*[]*T, error)
	Save(db *gorm.DB, data *T) (*T, error)
	UpdateOne(db *gorm.DB, data *T, selects ...string) (*T, error)
	Delete(db *gorm.DB, ids ...[]int) error
}

type CommonAdapter[T ICommonModel] struct{}

func (c *CommonAdapter[T]) GetOne(db *gorm.DB, where ...interface{}) (*T, error) {
	return _getOne[T](db, where)
}

func (c *CommonAdapter[T]) GetMulti(db *gorm.DB, where ...interface{}) (*[]*T, error) {
	return _getMulti[T](db, where)
}

func (c *CommonAdapter[T]) Save(db *gorm.DB, data *T) (*T, error) {
	return _save[T](db, data)
}

func (c *CommonAdapter[T]) UpdateOne(db *gorm.DB, data *T, selects ...string) (*T, error) {
	return _updateOne[T](db, data, selects...)
}

func (c *CommonAdapter[T]) Delete(db *gorm.DB, ids ...[]int) error {
	return _delete[T](db, ids...)
}

func _getOne[T ICommonModel](tx *gorm.DB, where ...interface{}) (*T, error) {
	var ret T
	err := tx.Model(&ret).Where(where).First(&ret).Error
	return &ret, err
}

func _getMulti[T ICommonModel](tx *gorm.DB, where ...interface{}) (*[]*T, error) {
	var ret []*T
	var name T
	err := tx.Model(&name).Where(where).Find(&ret).Error
	return &ret, err
}

func _save[T ICommonModel](tx *gorm.DB, data *T) (*T, error) {
	err := tx.Model(&data).Create(&data).Error
	return data, err
}

func _updateOne[T ICommonModel](tx *gorm.DB, data *T, selects ...string) (*T, error) {
	if selects == nil || len(selects) == 0 || selects[0] == "" {
		selects = []string{"*"}
	}
	err := tx.Model(&data).Select(selects).Updates(&data).Error
	return data, err
}

func _delete[T ICommonModel](tx *gorm.DB, ids ...[]int) (err error) {
	var model T
	err = tx.Model(&model).Delete(&model, ids).Error
	return err
}
