package model

import "time"

type Comment struct {
	ID        string `gorm:"primaryKey"`
	UserID    string
	PhotoID   string
	Message   string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Request
type CommentCreateRequest struct {
	Message string `json:"message" valid:"required~Message is required"`
}

type CommentUpdateRequest struct {
	Message string `json:"message" valid:"required~Message is required"`
}

// Response
type CommentCreateResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CommentUpdateResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	PhotoID   string    `json:"photo_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteCommentResponse struct {
	Message string `json:"message"`
}
