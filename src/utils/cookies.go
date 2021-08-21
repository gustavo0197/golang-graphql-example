package utils

import (
	"context"

	"github.com/gustavo0197/graphql/src/middlewares"
)

func GetCookiesWriter(ctx context.Context) *middlewares.Cookies {
	return ctx.Value("authorization").(*middlewares.Cookies)
}