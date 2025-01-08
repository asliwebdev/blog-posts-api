package service

import (
	"posts/models"
	"posts/repository"

	"github.com/google/uuid"
)

type PostService struct {
	postRepo *repository.PostRepo
}

func NewPostService(postRepo *repository.PostRepo) *PostService {
	return &PostService{postRepo: postRepo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.postRepo.CreatePost(post)
}

func (s *PostService) GetPostById(id uuid.UUID) (*models.Post, error) {
	return s.postRepo.GetPostById(id)
}

func (s *PostService) GetUserPosts(userId uuid.UUID) ([]models.Post, error) {
	return s.postRepo.GetUserPosts(userId)
}

func (s *PostService) GetFeedPosts(userId uuid.UUID) ([]models.Post, error) {
	return s.postRepo.GetFeedPosts(userId)
}

func (s *PostService) UpdatePost(post *models.Post) error {
	return s.postRepo.UpdatePost(post)
}

func (s *PostService) DeletePost(id uuid.UUID) error {
	return s.postRepo.DeletePost(id)
}
