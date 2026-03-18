package model

import (
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	if err := os.MkdirAll(".build/temp", 0o755); err != nil {
		return nil, err
	}
	db, err := gorm.Open(sqlite.Open(".build/temp/gorm.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	// 迁移 schema
	if err := db.AutoMigrate(&Product{}); err != nil {
		return nil, err
	}
	return db, nil
}
