package main

import (
	"net/http"

	"log"

	"github.com/99designs/gqlgen/handler"
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql/resolver"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
	"github.com/ryutah/go-graphql-photo-share-api/registry/local"
)

func main() {
	http.HandleFunc("/", handler.Playground("photo share api", "/query"))
	http.HandleFunc("/query", handler.GraphQL(graphql.NewExecutableSchema(graphql.Config{
		Resolvers: resolver.NewRoot(registry.NewProvider(local.InjectorConfig)),
	})))
	log.Println("start server on port 8080")
	log.Println(http.ListenAndServe(":8080", nil))
}
