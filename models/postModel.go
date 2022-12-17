package models

import (
	"github.com/jianrong/cvwo-be/initializers"
	"github.com/lib/pq"
)

type Post struct {
  Entry `gorm:"embedded"`
  Title string `gorm:"type:text" json:"title"`
  Tags pq.StringArray `gorm:"type:text[]" json:"tags"`

  // Upvotes []Rating `json:"upvotes"`
  // Downvotes []Rating `json:"downvotes"`

  Comments []Comment `json:"comments"`
}

func (post *Post) Save() (*Post, error) {
  err := initializers.DB.Create(&post).Error
  if err != nil {
    return &Post{}, err
  }
  return post, nil
}