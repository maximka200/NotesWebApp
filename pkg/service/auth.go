package service

import (
	todo "github.com/maximka200/NotesWebApp"
	"github.com/maximka200/NotesWebApp/pkg/repository"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(rrepo repository.Repository) *AuthService {
	return &AuthService{}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	var err error
	user.Password, err = generatePasswordHash(user.Password)
	if err != nil {
		logrus.Fatalf("error generate password: %s", err)
	}
	return s.repo.CreateUser(user)
}

// использую более безопасный bcrypt
func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}
