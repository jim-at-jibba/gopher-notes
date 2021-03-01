package service

import (
	"fmt"
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"github.com/jim-at-jibba/gopher-notes/pkg/repository"
)

type NoteService interface {
	CreateNote(note *model.Note) (*model.Note, error)
	CreateNoteFromRequest(model.NewNote) (*model.Note, error)
}

type noteService struct {
	repo repository.NoteRespository
}

func NewUserService(userRepo repository.NoteRespository) NoteService {
	return &noteService{repo: userRepo}
}

func (n *noteService) CreateNote(note *model.Note) (*model.Note, error) {

	fmt.Println("Service - creating")
	fmt.Printf("%+v", note)
	note, err := n.repo.CreateNote(note)

	if err != nil {
		fmt.Printf("Service error %v\n", err)
		return nil, err
	}
	return note, nil
}

func (n *noteService) CreateNoteFromRequest(creationRequest model.NewNote) (*model.Note, error) {
	//fmt.Println("Transforming", creationRequest)
	preparedNote, err := creationRequest.TransformToNote()
	//fmt.Printf("Prepared %+v", preparedNote)
	if err != nil {
		return nil, err
	}
	return n.CreateNote(preparedNote)
}
