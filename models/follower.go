package models

import (
	"time"

	"github.com/google/uuid"
)

type Follower struct {
	Id          uuid.UUID `json:"id"`
	FollowerId  uuid.UUID `json:"follower_id"`
	FollowingId uuid.UUID `json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type FollowRequest struct {
	FollowerId  uuid.UUID `json:"follower_id" binding:"required"`
	FollowingId uuid.UUID `json:"following_id" binding:"required"`
}
