package service

import (
	todo "WebApplication"
	"WebApplication/pkg/repository"
	"crypto/sha1"
	"fmt"
)

/**
 * Authorization service
 */

const salt = "hdkjahdkjakl312313819kldjaldk;l"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
