package models

import (
	"time"

	"github.com/jianrong/cvwo-be/initializers"
	"gorm.io/gorm"
)

type Entry struct {
  ID        uint           `gorm:"primaryKey;autoIncrement"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
	Content string `gorm:"type:text" json:"content"`
  Upvotes []Rating `json:"upvotes"`
  Downvotes []Rating `json:"downvotes"`
  UserID uint `json:"userId"`
  User User `json:"user"`
}

func (entry *Entry) Save() (*Entry, error) {
  err := initializers.DB.Create(&entry).Error
  if err != nil {
    return &Entry{}, err
  }
  return entry, nil
}