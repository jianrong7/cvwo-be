package models

import (
	"github.com/jianrong/cvwo-be/initializers"
	"gorm.io/gorm"
)

type Rating struct {
  gorm.Model
	// Value can tell if its an upvote or downvote. 1 is upvote, -1 is downvote.
	Value uint `json:"value"`

	// each rating belongs to a user.
	UserID uint `json:"userId"`
	User User `json:"user"`

	// each rating also belongs to an entry (post/comment)
	EntryID uint `json:"entryId"`
	// Entry Entry `json:"entry"`
	EntryType string `json:"entryType"`
}

func (rating *Rating) Save() (*Rating, error) {
  err := initializers.DB.Create(&rating).Error
  if err != nil {
    return &Rating{}, err
  }
  return rating, nil
}
