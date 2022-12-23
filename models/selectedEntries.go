package models

import "github.com/lib/pq"

type SelectedEntries struct {
	PostIds pq.Int64Array `gorm:"type:integer[]"`
	CommentIds pq.Int64Array `gorm:"type:integer[]"`
	UserId uint
	// RatingValue int
	// EntryType string
    // MaxTokens int `json:"maxTokens" binding:"required"`
    // Prompt string `json:"prompt" binding:"required"`
}