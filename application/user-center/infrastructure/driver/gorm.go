package driver

import (
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	gormDB *gorm.DB
	once   sync.Once
)

func GetGormDB() *gorm.DB {
	once.Do(func() {
		var err error
		gormDB, err = gorm.Open()
		if err != nil {
			log.Fatalf("Failed to initialize GORM: %v", err)
		}
	})
	return gormDB
}
