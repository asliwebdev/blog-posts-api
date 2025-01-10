package service

import (
	"fmt"
	"posts/pkg"
	"posts/repository"

	"posts/models"

	"github.com/google/uuid"
)

type UserService struct {
	userRepo     *repository.UserRepo
	followerRepo *repository.FollowerRepo
}

func NewUserService(userRepo *repository.UserRepo, followerRepo *repository.FollowerRepo) *UserService {
	return &UserService{userRepo: userRepo, followerRepo: followerRepo}
}

func (u *UserService) GetUserById(userId uuid.UUID) (*models.User, error) {
	user, err := u.userRepo.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	followerCount, followingCount, err := u.followerRepo.CountFollowersAndFollowing(userId)
	if err != nil {
		return nil, err
	}

	user.FollowerCount = followerCount
	user.FollowingCount = followingCount

	return user, nil
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *UserService) UpdateUser(user *models.UpdateUser) error {
	exists, err := u.userRepo.UserExists(user.Id)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	if user.Password != "" {
		hashedPassword, err := pkg.HashPassword(user.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = hashedPassword
	}

	return u.userRepo.UpdateUser(user)
}

func (u *UserService) DeleteUser(id uuid.UUID) error {
	exists, err := u.userRepo.UserExists(id)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	return u.userRepo.DeleteUser(id)
}
