package directives

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gustavo0197/graphql/src/utils"
)

func IsAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	cookies := utils.GetCookiesWriter(ctx)

	if !cookies.IsLoggedIn {
		return nil, errors.New("Access denied")
	}

	return next(ctx)
}