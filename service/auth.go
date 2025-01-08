package service

import (
	"errors"
	"posts/models"
	"posts/pkg"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrEmailExist         = errors.New("user with this email already exists")
	ErrUserNameExist      = errors.New("user with this username already exists")
)

func (u *UserService) Login(credentials *models.LoginRequest) (string, error) {
	user, err := u.userRepo.GetUserByEmail(credentials.Email)
	if err != nil {
		return "", err
	}

	if user == nil {
		return "", ErrInvalidCredentials
	}

	if err := pkg.ComparePasswords(user.Password, credentials.Password); err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := pkg.CreateToken(user.Username, user.Email, user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserService) SignUp(user *models.User) (string, error) {
	existingUser, err := u.userRepo.GetUserByEmailOrUsername(user.Email, user.Username)
	if err != nil {
		return "", err
	}
	if existingUser != nil {
		if existingUser.Email == user.Email {
			return "", ErrEmailExist
		}
		if existingUser.Username == user.Username {
			return "", ErrUserNameExist
		}
	}

	hashedPassword, err := pkg.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	if err := u.userRepo.CreateUser(user); err != nil {
		return "", err
	}

	token, err := pkg.CreateToken(user.Username, user.Email, user.Id)
	if err != nil {
		return "", err
	}

	return token, nil
}
