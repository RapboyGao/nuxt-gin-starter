// [框架]
// 本Package仅为保留gorm和sqlite库，没有实际作用
// [本文件可删除或者修改]

package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func Initialize() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	// 迁移 schema
	db.AutoMigrate(&Product{})
	return db, err
}
