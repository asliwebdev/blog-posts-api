package models

import (
	"time"

	"github.com/google/uuid"
)

type Follower struct {
	ID          uuid.UUID `json:"id"`
	FollowerID  uuid.UUID `json:"follower_id"`
	FollowingID uuid.UUID `json:"following_id"`
	CreatedAt   time.Time `json:"created_at"`
}
