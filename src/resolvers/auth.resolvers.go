package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/gustavo0197/graphql/src/model"
	"github.com/gustavo0197/graphql/src/services"
	"github.com/gustavo0197/graphql/src/utils"
)

func (r *queryResolver) Login(ctx context.Context, credentials model.UserCredentials) (bool, error) {
	authService := services.AuthService{}
	authService.MongoService = r.MongoService
	cookiesWriter := utils.GetCookiesWriter(ctx)

	token, error := authService.Login(&credentials)

	if error != nil || token == "" {
		return false, error
	}

	cookiesWriter.SetToken(token, time.Now().Add(time.Hour * 24 * 15))

	return true, error
}
