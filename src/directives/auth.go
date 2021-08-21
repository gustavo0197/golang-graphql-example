package directives

import (
	"context"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/gustavo0197/graphql/src/utils"
)

func IsAuth(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	cookies := utils.GetCookiesWriter(ctx)
	// TODO decode token and verify that is a valid token
	log.Println("isAuth directive: ", cookies.Token)

	return next(ctx)
}