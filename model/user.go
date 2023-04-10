package model

import (
	"time"
)

type User struct {
	ID           string `gorm:"primaryKey" `
	Username     string `gorm:"not null;unique;type:varchar(30)" `
	Email        string `gorm:"not null;unique;type:varchar(255)"`
	Password     string `gorm:"not null;type:varchar(255)"`
	Age          int    `gorm:"not null;size:2"`
	Photos       []Photo
	Comments     []Comment
	SocialMedias []SocialMedia
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type UserRegisterRequest struct {
	Email    string `json:"email" valid:"required~Email is required,email~Invalid email address"`
	Username string `json:"username" valid:"required~Username is required"`
	Password string `json:"password" valid:"required~Password is required,minstringlength(6)~Password atleast 6 characters"`
	Age      int    `json:"age" valid:"required~Age is required,range(8|99)~Age minimum is 8"`
}

type UserRegisterResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLoginRequest struct {
	Username string `json:"username" valid:"required~Username is required"`
	Password string `json:"password" valid:"required~Password is required"`
}
type UserLoginResponse struct {
	Token string `json:"token"`
}
