package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

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

func (r *mutationResolver) CreateUser(ctx context.Context, name string, email string, password string, usertype int) (*model.User, error) {
	rs, err := database.Handler.CreateUser(&name, &email, &password, &usertype)
	if err != nil {
		panic(err)
	} // returns error
	return rs, err
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, name *string, email *string, password *string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateIdentificationDocument(ctx context.Context, name *string) (*model.IdentificationDocument, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateIdentificationDocument(ctx context.Context, id string, name string) (*model.IdentificationDocument, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteIdentificationDocument(ctx context.Context, id string) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateClinicHistory(ctx context.Context, name string, documentURL *string, patientid int) (*model.ClinicHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateClinicHistory(ctx context.Context, id string, documentURL string) (*model.ClinicHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteClinicHistory(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePatient(ctx context.Context, documentID string, name string, lastname string, phoneNumber *string, sisID *string, genre string, birthDate time.Time, address string, district string, region string, identificationDocumentID int, userid int) (*model.Patient, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdatePatient(ctx context.Context, id string, documentID *string, name *string, lastname *string, phoneNumber *string, sisID *string, genre string, birthDate *time.Time, address *string, district *string, region *string, identificationDocumentID *int, userid *int, clinicHistoryID *int, documentIDURL *string) (*model.Patient, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePatient(ctx context.Context, id string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListUserTypes(ctx context.Context) ([]*model.UserType, error) {
	rs, err := database.Handler.ListUserTypes()
	if err != nil {
		panic(err)
	} // returns error
	return rs, err
}

func (r *queryResolver) ListUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListPatiens(ctx context.Context) ([]*model.Patient, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListClinicHistories(ctx context.Context) ([]*model.ClinicHistory, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListIdentificationDocuments(ctx context.Context) ([]*model.IdentificationDocument, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
