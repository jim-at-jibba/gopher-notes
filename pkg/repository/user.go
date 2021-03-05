package repository

import (
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *model.NewUser) (*model.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(user *model.NewUser) (*model.User, error) {
	var returnedUser *model.User
	return returnedUser, nil
}
