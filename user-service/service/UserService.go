package service

import (
	"errors"
	"user-service/model"
	"user-service/repo"
)

type UserService struct {
	Repo *repo.UserRepository
}

func NewUserService(repo *repo.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (service *UserService) Create(user *model.User) error {
	_, err := service.GetOne(user.Name)
	if err == nil {
		return errors.New("User with this name already exists")
	}
	return service.Repo.Create(user)
}

func (service *UserService) GetOne(name string) (*model.User, error) {
	return service.Repo.GetOne(name)
}
