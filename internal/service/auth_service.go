package service

import (
	"context"
	"errors"

	"github.com/lokicodess/CatalogX/internal/auth"
	"github.com/lokicodess/CatalogX/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type AuthService struct {
	UserRepo  repository.UserRepository
	JWTSecret string
}

func NewAuthService(userRepo repository.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{
		UserRepo:  userRepo,
		JWTSecret: jwtSecret,
	}
}

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	user, err := s.UserRepo.GetByEmail(ctx, email)
	if err != nil {
		return "", ErrInvalidCredentials
	}
	if user == nil {
		return "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", ErrInvalidCredentials
	}

	token, err := auth.GenerateToken(user.ID, user.Email, s.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
