package models

import (
	"github.com/jianrong/cvwo-be/initializers"
)

type Comment struct {
	Entry `gorm:"embedded"`
	Replies []*Comment `gorm:"many2many:replies"`
	PostID uint `gorm:"not null;" json:"postId"`
	Post Post `json:"post"`
}

func (comment *Comment) Save() (*Comment, error) {
  err := initializers.DB.Create(&comment).Error
  if err != nil {
    return &Comment{}, err
  }
  return comment, nil
}