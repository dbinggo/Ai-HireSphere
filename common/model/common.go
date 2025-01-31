package model

import (
	"context"
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

// 充血模型与schema互转
type ICommonEntity[T any, F any] interface {
	// Transform 转换为数据库schema结构
	Transform() F
	// From 从数据库schema结构转换为实体
	From(F) T
}

// 数据库接口适配器，提供最基本CURD能力
type IDBAdapter[T ICommonModel] interface {
	GetOne(ctx context.Context, db *gorm.DB, where interface{}, args ...interface{}) (*T, error)
	GetMulti(ctx context.Context, db *gorm.DB, where interface{}, args ...interface{}) (*[]*T, error)
	Save(ctx context.Context, db *gorm.DB, data *T) (*T, error)
	UpdateOne(ctx context.Context, db *gorm.DB, data *T, selects ...string) (*T, error)
	Delete(ctx context.Context, db *gorm.DB, ids ...int64) error
	List(ctx context.Context, tx *gorm.DB, limit int, offset int, where interface{}, args ...interface{}) (int64, []T, error)
}

type CommonAdapter[T ICommonModel] struct{}

func (c *CommonAdapter[T]) GetOne(ctx context.Context, db *gorm.DB, where interface{}, args ...interface{}) (*T, error) {
	return _getOne[T](ctx, db, where, args...)
}

func (c *CommonAdapter[T]) GetMulti(ctx context.Context, db *gorm.DB, where interface{}, args ...interface{}) (*[]*T, error) {
	return _getMulti[T](ctx, db, where, args)
}

func (c *CommonAdapter[T]) Save(ctx context.Context, db *gorm.DB, data *T) (*T, error) {
	return _save[T](ctx, db, data)
}

func (c *CommonAdapter[T]) UpdateOne(ctx context.Context, db *gorm.DB, data *T, selects ...string) (*T, error) {
	return _updateOne[T](ctx, db, data, selects...)
}

func (c *CommonAdapter[T]) Delete(ctx context.Context, db *gorm.DB, ids ...int64) error {
	return _delete[T](ctx, db, ids)
}
func (c *CommonAdapter[T]) List(ctx context.Context, tx *gorm.DB, limit int, offset int, where interface{}, args ...interface{}) (int64, []T, error) {
	return _list[T](ctx, tx, limit, offset, where, args...)
}

func _getOne[T ICommonModel](ctx context.Context, tx *gorm.DB, where interface{}, args ...interface{}) (*T, error) {
	var ret T
	err := tx.WithContext(ctx).Model(&ret).Where(where, args...).First(&ret).Error
	return &ret, err
}

func _getMulti[T ICommonModel](ctx context.Context, tx *gorm.DB, where interface{}, args ...interface{}) (*[]*T, error) {
	var ret []*T
	var name T
	err := tx.WithContext(ctx).Model(&name).Where(where, args...).Find(&ret).Error
	return &ret, err
}

func _save[T ICommonModel](ctx context.Context, tx *gorm.DB, data *T) (*T, error) {
	err := tx.WithContext(ctx).Model(&data).Create(&data).Error
	return data, err
}

func _updateOne[T ICommonModel](ctx context.Context, tx *gorm.DB, data *T, selects ...string) (*T, error) {
	if selects == nil || len(selects) == 0 || selects[0] == "" {
		selects = []string{"*"}
	}
	err := tx.WithContext(ctx).Model(&data).Select(selects).Updates(&data).Error
	return data, err
}

func _delete[T ICommonModel](ctx context.Context, tx *gorm.DB, ids ...[]int64) (err error) {
	var model T
	err = tx.WithContext(ctx).Model(&model).Delete(&model, ids).Error
	return err
}

func _list[T ICommonModel](ctx context.Context, tx *gorm.DB, limit int, offset int, where interface{}, args ...interface{}) (int64, []T, error) {
	var ret []T
	var name T
	var count int64
	err := tx.WithContext(ctx).Model(&name).Where(where, args...).Limit(limit).Offset(offset).Count(&count).Find(&ret).Error
	return count, ret, err
}
