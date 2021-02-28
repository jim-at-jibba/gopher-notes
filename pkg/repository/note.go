package repository

import (
	"fmt"
	"github.com/gofrs/uuid"
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"github.com/jmoiron/sqlx"
	"log"
)

type NoteRespository interface {
	CreateNote(note *model.Note) (*model.Note, error)
}

type noteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRespository {
	return &noteRepository{
		db: db,
	}
}

func (n *noteRepository) CreateNote(note *model.Note) (*model.Note, error) {
	if note.ID == "" {
		note.ID = uuid.Must(uuid.NewV4()).String()
	}
	newNoteStatement := `INSERT INTO "notes" (title, text, user_id) VALUES ($1, $2, $3)`

	err := n.db.Get(newNoteStatement, note.Title, note.Title, 1)

	if err != nil {
		log.Printf("error creating the note: %v", err)
		return nil, err
	}
	fmt.Printf(
		"Created note: %v",
		note,
	)
	return note, nil

}
