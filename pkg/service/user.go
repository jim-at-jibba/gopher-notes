package service

import (
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"github.com/jim-at-jibba/gopher-notes/pkg/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(user *model.DBUser) (*model.User, error)
	CreateUserFromRequest(creationRequest model.NewUser) (*model.User, error)
	GetUserIdByUsername(username string) (string, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		repo: userRepo,
	}
}

func (u *userService) CreateUser(user *model.DBUser) (*model.User, error) {

	returnedUser, err := u.repo.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return returnedUser, nil
}

func (u *userService) CreateUserFromRequest(creationRequest model.NewUser) (*model.User, error) {
	prepparedUser, err := creationRequest.TransformToUser()
	hash, err := HashPassword(creationRequest.Password)
	if err != nil {
		return nil, err
	}

	DbUser := &model.DBUser{
		ID:       prepparedUser.ID,
		Username: prepparedUser.Username,
		Password: hash,
	}

	return u.CreateUser(DbUser)
}

//HashPassword hashes given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPassword hash compares raw password with it's hashed values
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u *userService) GetUserIdByUsername(username string) (string, error) {

	returnedUserId, err := u.repo.GetUserIdByUsername(username)

	if err != nil {
		return "", err
	}

	return returnedUserId, nil
}
