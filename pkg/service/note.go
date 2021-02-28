package service

import (
	"github.com/jim-at-jibba/gopher-notes/pkg/repository"
)

type NoteService interface{}

type noteService struct {
	repo repository.NoteRespository
}

func NewUserService(userRepo repository.NoteRespository) NoteService {
	return &noteService{repo: userRepo}
}
