package service

import (
	todo "github.com/maximka200/NotesWebApp"
	"github.com/maximka200/NotesWebApp/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

const solt = "aaaasssssdddd123123kasdkaksdk"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(rrepo repository.Repository) *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	return s.repo.CreateUser(user)
}

// использую более безопасный bcrypt
func (s *AuthService) generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
