package models

import (
	"github.com/jianrong/cvwo-be/initializers"
	"gorm.io/gorm"
)

type Comment struct {
  gorm.Model
	Content string `json:"content"`
	// Replies []*Comment `gorm:"many2many:replies"`
	PostID uint `gorm:"not null;" json:"postId"`
	Post Post `json:"post"`
	UserID uint `json:"userId"`
	User User `json:"user"`
}

func (comment *Comment) Save() (*Comment, error) {
  err := initializers.DB.Create(&comment).Error
  if err != nil {
    return &Comment{}, err
  }
  return comment, nil
}