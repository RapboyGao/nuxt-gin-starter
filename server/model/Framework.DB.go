package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	// 迁移 schema
	db.AutoMigrate(&Product{})
	return db, err
}
