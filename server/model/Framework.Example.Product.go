// [框架]
// 本Package仅为保留gorm和sqlite库，没有实际作用
// [本文件可删除或者修改]

package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}
