package models

import (
	"gorm.io/gorm"
)

type Post struct {
  gorm.Model
  Title string
	Body string
  // Tags pq.StringArray `gorm:type:text[]"`
  UserId uint
  User User `gorm:"foreignKey:UserId"`
  Upvotes uint
  Downvotes uint
  Comments []Comment
}