package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	ActorID   uuid.UUID `json:"actor_id"`
	PostID    uuid.UUID `json:"post_id,omitempty"`
	CommentID uuid.UUID `json:"comment_id,omitempty"`
	FollowID  uuid.UUID `json:"follow_id,omitempty"`
	Type      string    `json:"type"`
	Message   string    `json:"message,omitempty"`
	Read      bool      `json:"read"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
