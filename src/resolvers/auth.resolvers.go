package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gustavo0197/graphql/src/model"
	"github.com/gustavo0197/graphql/src/services"
)

func (r *queryResolver) Login(ctx context.Context, credentials model.UserCredentials) (string, error) {
	authService := services.AuthService{}
	authService.MongoService = r.MongoService

	return authService.Login(&credentials)
}
