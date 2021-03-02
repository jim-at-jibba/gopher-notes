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
	ListNotes(userId string) ([]*model.Note, error)
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

func (n *noteRepository) ListNotes(userId string) ([]*model.Note, error) {

	var dbnotes []*model.DBNote
	var notes []*model.Note

	if err := n.db.Select(&dbnotes, `SELECT id, title, text, user_id, created_at FROM "notes" WHERE user_id=$1`, userId); err != nil {
		log.Printf("error getting the notes: %v", err)
		return []*model.Note{}, err
	}

	for _, note := range dbnotes {
		fmt.Printf("Note %v \n", note)
		var nt *model.Note

		user := &model.User{
			ID: note.UserID,
		}

		nt = &model.Note{
			ID:        note.ID,
			Title:     note.Title,
			Text:      note.Text,
			User:      user,
			CreatedAt: note.CreatedAt,
		}
		notes = append(notes, nt)
	}
	return notes, nil
}
