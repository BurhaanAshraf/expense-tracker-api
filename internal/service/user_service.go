package service

import (
	"context"
	"errors"

	"github.com/BurhaanAshraf/finance-api/internal/auth"
	"github.com/BurhaanAshraf/finance-api/internal/model"
	"github.com/BurhaanAshraf/finance-api/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) Register(ctx context.Context, name, email, password string) (*model.User, error) {
	existingUser, err := s.userRepository.FindByEmail(ctx, email)

	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Name:         name,
		Email:        email,
		PasswordHash: hashedPassword,
	}

	err = s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) Login(ctx context.Context, email, password string, secret string) (string, error) {
	user, err := s.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = auth.CheckPassword(password, user.PasswordHash)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateToken(user.ID, secret)
	if err != nil {
		return "", err
	}

	return token, nil
}
