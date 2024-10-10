package services

import (
	"anki-project/models"
	"anki-project/repository"
)

type UserService interface {
	Register(registerData models.DBUser) (models.DBUser, error)
	Login(login, password string) (models.DBUser, error)
	Edit(userId int, user models.EditUser) (models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Register(user models.DBUser) (models.DBUser, error) {
	return s.repo.Register(user)
}

func (s *userService) Login(login, password string) (models.DBUser, error) {
	return s.repo.Login(login, password)
}

func (s *userService) Edit(userId int, user models.EditUser) (models.User, error) {
	return s.repo.Edit(userId, user)
}
