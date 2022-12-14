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

  UserID uint `json:"userId"`
  User User `json:"user"`

  // RatingID uint `json:"ratingId"`
  // Rating Rating `json:"rating"`

  Comments []Comment `json:"comments"`

  Upvotes uint `json:"upvotes"`
  Downvotes uint `json:"downvotes"`
}

func (post *Post) Save() (*Post, error) {
  err := initializers.DB.Create(&post).Error
  if err != nil {
    return &Post{}, err
  }
  return post, nil
}