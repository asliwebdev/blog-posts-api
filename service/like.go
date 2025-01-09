package service

import (
	"posts/models"
	"posts/repository"

	"github.com/google/uuid"
)

type LikeService struct {
	repo *repository.LikeRepo
}

func NewLikeService(repo *repository.LikeRepo) *LikeService {
	return &LikeService{repo: repo}
}

func (s *LikeService) ToggleLike(userId, postId, commentId uuid.UUID) error {
	likeExists, err := s.repo.CheckIfLiked(userId, postId, commentId)
	if err != nil {
		return err
	}

	if likeExists {
		return s.repo.DeleteLike(userId, postId, commentId)
	}

	newLike := &models.Like{
		Id:        uuid.New(),
		UserId:    userId,
		PostId:    postId,
		CommentId: commentId,
	}

	return s.repo.AddLike(newLike)
}

func (s *LikeService) GetLikedUsers(postId, commentId uuid.UUID) ([]models.UserResponse, error) {
	return s.repo.GetLikedUsers(postId, commentId)
}
