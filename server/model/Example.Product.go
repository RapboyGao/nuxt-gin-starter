// [框架]
// 本Package仅为保留gorm和sqlite库，没有实际作用
// [本文件可删除或者修改]

package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `gorm:"size:128;not null"`
	Price float64 `gorm:"not null"`
	Code  string  `gorm:"size:64"`
	Level string  `gorm:"size:16;not null;default:standard"`
}
