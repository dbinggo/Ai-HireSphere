package test

import (
	"Ai-HireSphere/common/model"
	"gorm.io/gorm"
	"testing"
)

func TestGorm(t *testing.T) {

	db := gorm.DB{}

	db.Transaction(func(tx *gorm.DB) error {
		ret := &model.Example{}
		model.NewService().GetOne(tx, "阿爸阿爸", ret)
		return nil
	})
}
