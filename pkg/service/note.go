package service

import (
	"fmt"
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
	"github.com/jim-at-jibba/gopher-notes/pkg/repository"
)

type NoteService interface {
	CreateNote(note *model.Note) (*model.Note, error)
}

type noteService struct {
	repo repository.NoteRespository
}

func NewUserService(userRepo repository.NoteRespository) NoteService {
	return &noteService{repo: userRepo}
}

func (n *noteService) CreateNote(note *model.Note) (*model.Note, error) {

	fmt.Println("Service - creating")
	note, err := n.repo.CreateNote(note)

	if err != nil {
		fmt.Println("Service error")
		return nil, err
	}
	return note, nil
}
