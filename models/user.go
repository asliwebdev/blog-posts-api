package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID `json:"id"`
	Username       string    `json:"username" binding:"required"`
	Email          string    `json:"email" binding:"required,email"`
	Password       string    `json:"password" binding:"required,min=4"`
	FollowingCount int       `json:"following_count"`
	FollowerCount  int       `json:"follower_count"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type LoginRequest struct {
	Id       uuid.UUID `json:"id"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password" binding:"required,min=4"`
}

type UserResponse struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
}

type UpdateUser struct {
	Id       uuid.UUID `json:"id"`
	Username string    `json:"username" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Password string    `json:"password"`
}
