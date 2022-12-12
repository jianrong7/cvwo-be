package models

import (
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Post struct {
  gorm.Model
  Title string `gorm:"type:text" json:"title"`
	Content string `gorm:"type:text" json:"content"`
  Tags pq.StringArray `gorm:"type:text[]" json:"tags"`
  UserId uint `json:"userId"`
  // User User `gorm:"foreignKey:UserId" json:"user"`
  Upvotes uint `gorm:"size:255" json:"upvotes"`
  Downvotes uint `gorm:"size:255" json:"downvotes"`
  Comments []Comment `json:"comments"`
}

func (post *Post) Save() (*Post, error) {
  err := initializers.DB.Create(&post).Error
  if err != nil {
    return &Post{}, err
  }
  return post, nil
}