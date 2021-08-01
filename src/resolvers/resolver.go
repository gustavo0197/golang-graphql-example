package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/gustavo0197/graphql/src/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	services.MongoService
}
