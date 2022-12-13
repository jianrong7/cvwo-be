package models

import "gorm.io/gorm"

type Comment struct {
  gorm.Model
	Content string
	Replies []*Comment `gorm:"many2many:replies"`
	PostId uint
	Post Post `gorm:"foreignKey:PostId"`
	UserId uint
	User User
}