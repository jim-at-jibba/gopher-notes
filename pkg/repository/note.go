package repository

import (
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

	var dbnote model.DBNote

	if note.ID == "" {
		note.ID = uuid.Must(uuid.NewV4()).String()
	}
	newNoteStatement := `INSERT INTO notes (id, title, text, user_id) VALUES ($1, $2, $3, $4) RETURNING id, title, text, user_id`

	//_, err := n.db.Exec(newNoteStatement, note.ID, note.Title, note.Text, "b823d4e2-7993-4c3c-be3a-bc3f44d1cec2")
	err := n.db.Get(&dbnote, newNoteStatement, note.ID, note.Title, note.Text, "b823d4e2-7993-4c3c-be3a-bc3f44d1cec2")
	if err != nil {
		log.Printf("error creating the note: %v", err)
		return nil, err
	}

	return note, nil

}
