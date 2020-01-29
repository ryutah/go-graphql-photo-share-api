package resolver

import (
	"github.com/99designs/gqlgen/handler"
	"github.com/labstack/echo/v4"
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql/dataloader"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
	"github.com/ryutah/go-graphql-photo-share-api/registry/local"
)

func NewHandler(p *registry.Provider) *echo.Echo {
	e := echo.New()

	e.GET("/", echo.WrapHandler(handler.Playground("photo share api", "/query")))
	e.POST(
		"/query",
		echo.WrapHandler(handler.GraphQL(
			graphql.NewExecutableSchema(graphql.Config{
				Resolvers: newRoot(registry.NewProvider(local.InjectorConfig)),
			})),
		),
		dataloader.Middleware(p),
	)

	return e
}
