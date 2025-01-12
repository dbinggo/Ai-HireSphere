package dataaccess

import (
	idataaccess "Ai-HireSphere/application/user-center/domain/irepository/idata_access"
	"gorm.io/gorm"
)

// GormOpts ...
type GormOpts struct {
	db *gorm.DB
}

var _ idataaccess.IDataAccess[GormOpts] = (*GormOpts)(nil)

func NewGormOpts(db *gorm.DB) *GormOpts {
	return &GormOpts{db: db}
}

func (o *GormOpts) WithTx(tx *gorm.DB) *GormOpts {
	o.db = tx
	return o
}

// 开启事务
func (o *GormOpts) Transaction(fc func(tx *GormOpts) error) {
	o.db.Transaction(func(db *gorm.DB) error {
		// 注入事务到 GormOpts中
		tx := NewGormOpts(db)
		// 执行业务逻辑
		return fc(tx)
	})
}
