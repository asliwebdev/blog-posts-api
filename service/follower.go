package service

import (
	"posts/models"
	"posts/repository"

	"github.com/google/uuid"
)

type FollowerService struct {
	repo *repository.FollowerRepo
}

func NewFollowerService(repo *repository.FollowerRepo) *FollowerService {
	return &FollowerService{repo: repo}
}

func (s *FollowerService) AddFollower(follower *models.FollowRequest) error {
	return s.repo.AddFollower(follower)
}

func (s *FollowerService) RemoveFollower(followerId, followingId uuid.UUID) error {
	return s.repo.RemoveFollower(followerId, followingId)
}

func (s *FollowerService) GetFollowers(userId uuid.UUID) ([]models.UserResponse, error) {
	return s.repo.GetFollowers(userId)
}

func (s *FollowerService) GetFollowing(userId uuid.UUID) ([]models.UserResponse, error) {
	return s.repo.GetFollowing(userId)
}
