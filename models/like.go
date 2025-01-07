package models

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id,omitempty"`
	CommentID uuid.UUID `json:"comment_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}
