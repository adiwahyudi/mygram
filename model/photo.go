package model

import "time"

type Photo struct {
	ID        string `gorm:"primaryKey"`
	Title     string `gorm:"not null;type:varchar(100)"`
	Caption   string `gorm:"not null;type:varchar(255)"`
	PhotoURL  string `gorm:"not null;type:varchar(255);column:photo_url"`
	UserID    string
	Comments  []Comment
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Request
type PhotoCreateRequest struct {
	Title    string `json:"title" valid:"required~Title is required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" valid:"required~Photo URL is required"`
}

type PhotoUpdateRequest struct {
	Title    string `json:"title" valid:"required~Title is required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" valid:"required~Photo URL is required"`
}

// Response
type PhotoCreateResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
}

type PhotoUpdateResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PhotoResponse struct {
	ID        string                   `json:"id"`
	UserID    string                   `json:"user_id"`
	Title     string                   `json:"title"`
	Caption   string                   `json:"caption"`
	PhotoURL  string                   `json:"photo_url"`
	Comments  []CommentInPhotoResponse `json:"comments"`
	CreatedAt time.Time                `json:"created_at"`
	UpdatedAt time.Time                `json:"updated_at"`
}

type CommentInPhotoResponse struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeletePhotoResponse struct {
	Message string `json:"message"`
}
