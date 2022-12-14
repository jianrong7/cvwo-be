package models

import (
	"gorm.io/gorm"
)

type Rating struct {
  gorm.Model
	Value uint `json:"value"`
	UserId uint `json:"userId"`
	User User `json:"user"`
	EntryId uint `json:"entryId"`
}
