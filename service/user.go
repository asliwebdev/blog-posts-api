package service

import (
	"fmt"
	"posts/repository"

	"posts/models"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService(userRepo *repository.UserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) CreateUser(user models.User) error {
	existingUser, err := u.userRepo.GetUserByEmailOrUsername(user.Email, user.Username)
	if err != nil {
		return err
	}
	if existingUser != nil {
		if existingUser.Email == user.Email {
			return fmt.Errorf("user with email %s already exists", user.Email)
		}
		if existingUser.Username == user.Username {
			return fmt.Errorf("user with username %s already exists", user.Username)
		}
	}

	return u.userRepo.CreateUser(user)
}

func (u *UserService) GetUserById(userId string) (*models.User, error) {
	return u.userRepo.GetUserById(userId)
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return u.userRepo.GetAllUsers()
}

func (u *UserService) UpdateUser(id string, user models.User) error {
	exists, err := u.userRepo.UserExists(id)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	return u.userRepo.UpdateUser(id, user)
}

func (u *UserService) DeleteUser(id string) error {
	exists, err := u.userRepo.UserExists(id)
	if err != nil {
		return fmt.Errorf("failed to check user existence: %w", err)
	}
	if !exists {
		return fmt.Errorf("user not found")
	}

	return u.userRepo.DeleteUser(id)
}
