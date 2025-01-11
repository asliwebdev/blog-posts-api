package models

import (
	"time"

	"github.com/google/uuid"
)

type Like struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	PostId    uuid.UUID `json:"post_id,omitempty"`
	CommentId uuid.UUID `json:"comment_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
}

type ToggleLikeRequest struct {
	PostId    uuid.UUID `json:"post_id,omitempty" binding:"omitempty,uuid"`
	CommentId uuid.UUID `json:"comment_id,omitempty" binding:"omitempty,uuid"`
}
