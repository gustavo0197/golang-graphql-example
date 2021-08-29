package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/gustavo0197/graphql/src/model"
	"github.com/gustavo0197/graphql/src/services"
	"github.com/gustavo0197/graphql/src/utils"
)

func (r *mutationResolver) User(ctx context.Context, data model.NewUser) (*model.User, error) {
	userService := services.UserService{}
	userService.MongoService = r.MongoService

	return userService.CreateUser(data)
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	userService := services.UserService{}
	userService.MongoService = r.MongoService
	cookies := utils.GetCookiesWriter(ctx)

	return userService.GetUser(cookies.UserId)
}
