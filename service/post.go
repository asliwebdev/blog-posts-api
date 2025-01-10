package service

import (
	"posts/models"
	"posts/repository"

	"github.com/google/uuid"
)

type PostService struct {
	postRepo    *repository.PostRepo
	likeRepo    *repository.LikeRepo
	commentRepo *repository.CommentRepo
}

func NewPostService(postRepo *repository.PostRepo, likeRepo *repository.LikeRepo, commentRepo *repository.CommentRepo) *PostService {
	return &PostService{postRepo: postRepo, likeRepo: likeRepo, commentRepo: commentRepo}
}

func (s *PostService) CreatePost(post *models.Post) error {
	return s.postRepo.CreatePost(post)
}

func (s *PostService) GetPostById(postId uuid.UUID) (*models.Post, error) {
	post, err := s.postRepo.GetPostById(postId)
	if err != nil {
		return nil, err
	}

	likesCount, err := s.likeRepo.CountPostLikes(postId)
	if err != nil {
		return nil, err
	}

	commentsCount, err := s.commentRepo.CountComments(postId)
	if err != nil {
		return nil, err
	}

	post.LikesCount = likesCount
	post.CommentsCount = commentsCount

	return post, nil
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
