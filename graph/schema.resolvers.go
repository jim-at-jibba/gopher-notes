package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/jim-at-jibba/gopher-notes/graph/generated"
	"github.com/jim-at-jibba/gopher-notes/pkg/jwt"
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
)

func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.Note, error) {
	note, err := r.NoteService.CreateNoteFromRequest(input)

	if err != nil {
		return nil, err
	}

	return note, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user, err := r.UserService.CreateUserFromRequest(input)

	token, err := jwt.GenerateToken(input.Username, user.ID)
	fmt.Println("Token %v:", token)
	if err != nil {
		return nil, err
	}

	return user, err
}

func (r *queryResolver) Notes(ctx context.Context, userID string) ([]*model.Note, error) {
	notes, err := r.NoteService.ListNotes(userID)

	if err != nil {
		return nil, err
	}

	return notes, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
