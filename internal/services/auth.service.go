package services

import (
	"errors"
	"strconv"

	"github.com/Chetan7595/task-manager/internal/models"
	"github.com/Chetan7595/task-manager/internal/repository"
	"github.com/Chetan7595/task-manager/pkg/utils"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(name, email, password string) error {

	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hashed,
		Role:     "USER",
	}

	return s.userRepo.Create(user)
}

func (s *AuthService) Login(email, password string) (string, error) {

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = utils.CheckPassword(password, user.Password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(
		strconv.FormatInt(user.ID, 10),
		user.Role,
	)
	if err != nil {
		return "", err
	}

	return token, nil
}
