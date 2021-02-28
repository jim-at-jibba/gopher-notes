package graph

import "github.com/jim-at-jibba/gopher-notes/pkg/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	NoteService service.NoteService
}
