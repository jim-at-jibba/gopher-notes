package repository

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"github.com/jmoiron/sqlx"
	"log"
)

type UserRepository interface {
	CreateUser(user *model.DBUser) (*model.User, error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUser(dbUser *model.DBUser) (*model.User, error) {

	if dbUser.ID == "" {
		dbUser.ID = uuid.Must(uuid.NewV4()).String()
	}

	newUserStatement := `INSERT INTO users (id, username, password) VALUES ($1, $2, $3) RETURNING id, username`

	_, err := u.db.Exec(newUserStatement, dbUser.ID, dbUser.Username, dbUser.Password)

	fmt.Printf("%+v", dbUser)
	if err != nil {
		log.Printf("error creating the user: %v", err)
		return nil, err
	}

	return &model.User{
		ID:       dbUser.ID,
		Username: dbUser.Username,
	}, nil
}
