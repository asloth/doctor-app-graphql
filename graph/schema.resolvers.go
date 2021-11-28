package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/asloth/doctor-app-graphql/database"
	"github.com/asloth/doctor-app-graphql/graph/generated"
	"github.com/asloth/doctor-app-graphql/graph/model"
)

func (r *mutationResolver) CreateUserType(ctx context.Context, input model.NewUserType) (*model.UserType, error) {
	rs, err := database.Handler.CreateUserType(&input.Name)
	if err != nil {
		panic(err)
	} // returns error
	return rs, err
}

func (r *mutationResolver) UpdateUserType(ctx context.Context, id string, input model.NewUserType) (*model.UserType, error) {
	rs, err := database.Handler.UpdateUserType(&id, &input.Name)
	if err != nil {
		panic(err)
	} // returns error
	return rs, err
}

func (r *mutationResolver) DeleteUserType(ctx context.Context, id string) (*model.UserType, error) {
	rs, err := database.Handler.DeleteUserType(&id)
	if err != nil {
		panic(err)
	} // returns error
	return rs, err
}

func (r *queryResolver) ListUserTypes(ctx context.Context) ([]*model.UserType, error) {
	rs, err := database.Handler.ListUserTypes()
	if err != nil {
		panic(err)
	} // returns error
	return rs, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
