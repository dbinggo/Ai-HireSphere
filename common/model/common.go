package model

import (
	"gorm.io/gorm"
	"time"
)

// 本类文件提供给最基础的增删改查语句，便于扩展

type CommonModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type ICommonModel interface {
	TableName() string
}

func _getOne[T ICommonModel](tx *gorm.DB, where interface{}) (T, error) {
	var ret T
	err := tx.Model(&ret).Where(where).First(&ret).Error
	return ret, err
}

func _getMulti[T ICommonModel](tx *gorm.DB, tableName string, where interface{}) ([]T, error) {
	var ret []T
	var name T
	err := tx.Model(&name).Where(where).Find(&ret).Error
	return ret, err
}

func _create[T ICommonModel](tx *gorm.DB, data T) error {
	err := tx.Model(&data).Create(&data).Error
	return err
}

func _updateOne[T ICommonModel](tx *gorm.DB, data T, selects ...string) (T, error) {
	if selects == nil {
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
