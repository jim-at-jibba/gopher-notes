package repository

import (
	"fmt"
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"github.com/jmoiron/sqlx"
)

type NoteRespository interface {
	CreateNote(note *model.Note)
}

type noteRepository struct {
	db *sqlx.DB
}

func NewNoteRepository(db *sqlx.DB) NoteRespository {
	return &noteRepository{
		db: db,
	}
}

func (s *noteRepository) CreateNote(note *model.Note) {
	fmt.Printf("Creating a note: %v", note)
}
