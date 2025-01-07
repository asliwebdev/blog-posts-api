package models

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID              uuid.UUID `json:"id"`
	PostID          uuid.UUID `json:"post_id"`
	UserID          uuid.UUID `json:"user_id"`
	ParentCommentID uuid.UUID `json:"parent_comment_id,omitempty"`
	Content         string    `json:"content"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
