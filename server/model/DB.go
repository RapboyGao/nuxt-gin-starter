package model

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	if err := os.MkdirAll(".build/temp", 0o755); err != nil {
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(".build/temp/gorm.db"), &gorm.Config{})
	// 迁移 schema
	db.AutoMigrate(&Product{})
	return db, err
}
