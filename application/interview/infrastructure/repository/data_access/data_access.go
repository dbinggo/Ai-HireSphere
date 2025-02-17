package data_access

import (
	"gorm.io/gorm"
)

type GormOpts struct {
	db *gorm.DB
}

func NewGormOpts(db *gorm.DB) *GormOpts {
	return &GormOpts{
		db: db,
	}
}
