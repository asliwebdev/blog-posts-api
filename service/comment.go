package service

import (
	"posts/models"
	"posts/repository"

	"github.com/google/uuid"
)

type CommentService struct {
	repo *repository.CommentRepo
}

func NewCommentService(repo *repository.CommentRepo) *CommentService {
	return &CommentService{repo: repo}
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	return s.repo.CreateComment(comment)
}

func (s *CommentService) GetCommentsByPostId(postId uuid.UUID) ([]models.Comment, error) {
	return s.repo.GetCommentsByPostId(postId)
}

func (s *CommentService) UpdateComment(comment *models.UpdateComment) error {
	return s.repo.UpdateComment(comment)
}

func (s *CommentService) DeleteComment(commentId, userId uuid.UUID) error {
	return s.repo.DeleteComment(commentId, userId)
}
