package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/jim-at-jibba/gopher-notes/graph/generated"
	model1 "github.com/jim-at-jibba/gopher-notes/graph/model"
	"github.com/jim-at-jibba/gopher-notes/pkg/model"
)

func (r *mutationResolver) CreateNote(ctx context.Context, input model.NewNote) (*model.Note, error) {
	note, err := r.NoteService.CreateNoteFromRequest(input)

	if err != nil {
		return nil, err
	}

	return note, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model1.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Notes(ctx context.Context, userID string) ([]*model.Note, error) {
	notes, err := r.NoteService.ListNotes(userID)

	if err != nil {
		return nil, err
	}

	return notes, err
}

func (r *userResolver) Username(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) CreateUsername(ctx context.Context, input model1.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
