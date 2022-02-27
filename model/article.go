package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title    string
	Content  string
	Author   string
	Creator  uint   `gorm:"creator"`
	Modifier uint   `gorm:"modifier"`
	User     User
	UserID   uint
}
