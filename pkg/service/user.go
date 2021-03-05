package service

import (
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"github.com/jim-at-jibba/gopher-notes/pkg/repository"
)

type UserService interface {
	CreateUser(user *model.NewUser) (*model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		repo: userRepo,
	}
}

func (u *userService) CreateUser(user *model.NewUser) (*model.User, error) {
	returnedUser, err := u.repo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return returnedUser, nil
}
