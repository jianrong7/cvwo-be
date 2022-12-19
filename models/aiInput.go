package models

type AiInput struct {
    MaxTokens int `json:"maxTokens" binding:"required"`
    Prompt string `json:"prompt" binding:"required"`
}