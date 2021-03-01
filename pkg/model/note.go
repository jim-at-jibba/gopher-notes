package model

import (
	"database/sql"
	"errors"
)

type NewNote struct {
	Text   string `json:"text"`
	Title  string `json:"title"`
	UserID string `json:"userId"`
}

type Note struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Title     string `json:"title"`
	User      *User  `json:"user"`
	CreatedAt string `json:"createdAt"`
}

type DBNote struct {
	ID        string         `json:"id" db:"id"`
	Title     string         `json:"title" db:"title"`
	Text      string         `json:"text" db:"text"`
	UserID    string         `json:"userId" db:"user_id"`
	CreatedAt string         `json:"createdAt" db:"created_at"`
	UpdatedAt sql.NullString `json:"updatedAt" db:"updated_at"`
}

func (creationRequest *NewNote) IsValid() bool {
	return true // Could add some validation logic here
}

// Transforms a NewNote-Request into an Note-Object.
// Returns the new Article-Object or an error, if the CreateArticle-Request is not valid.
func (creationRequest *NewNote) TransformToNote() (*Note, error) {
	if !creationRequest.IsValid() {
		return nil, errors.New("article object is not valid")
	}

	// Hopefully this will come from JWT
	user := &User{
		ID: creationRequest.UserID,
	}

	//fmt.Printf("In transform: %+v\n", *user)
	return &Note{
		Title: creationRequest.Title,
		Text:  creationRequest.Text,
		User:  user,
	}, nil
}
