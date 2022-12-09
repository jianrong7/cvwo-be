package models

import "gorm.io/gorm"

type Comment struct {
  gorm.Model
	Body string
	// Replies []Comment
	PostId uint
	Post Post `gorm:"foreignKey:PostId"`
}