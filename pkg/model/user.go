package model

import "errors"

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
}

type NewUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DBUser struct {
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

func (creationRequest *NewUser) IsValid() bool {
	return true // Could add some validation logic here
}
func (creationRequest *NewUser) TransformToUser() (*User, error) {
	if !creationRequest.IsValid() {
		return nil, errors.New("user object not valid")
	}

	return &User{
		Username: creationRequest.Username,
	}, nil
}
