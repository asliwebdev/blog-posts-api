package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	Id              uuid.UUID    `json:"id"`
	PostId          uuid.UUID    `json:"post_id" binding:"required"`
	UserId          uuid.UUID    `json:"user_id"`
	ParentCommentId uuid.UUID    `json:"parent_comment_id,omitempty"`
	Content         string       `json:"content" binding:"required"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
	User            UserResponse `json:"user"`
}

type UpdateComment struct {
	Id      uuid.UUID `json:"id" binding:"required,uuid"`
	Content string    `json:"content" binding:"required"`
	UserId  uuid.UUID `json:"user_id"`
}
