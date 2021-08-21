package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/gustavo0197/graphql/src/directives"
	"github.com/gustavo0197/graphql/src/generated"
	"github.com/gustavo0197/graphql/src/middlewares"
	"github.com/gustavo0197/graphql/src/resolvers"
	"github.com/joho/godotenv"
)

const defaultPort = "8080"

// GraphQL handler
func graphqlHandler() gin.HandlerFunc {
	// Connect to MongoDB
	resolver := resolvers.Resolver{}
	resolver.MongoService.Connect()
	resolver.MongoService.CreateCollections()
	config := generated.Config{}
	config.Resolvers = &resolver

	config.Directives.IsAuth = directives.IsAuth

	h := handler.NewDefaultServer(generated.NewExecutableSchema(config))

	return func (c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = defaultPort
	}

	r := gin.Default()

	r.Use(middlewares.Auth())

	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
	r.Run("localhost:" + port)
}
