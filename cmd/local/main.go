package main

import (
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql/resolver"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
	"github.com/ryutah/go-graphql-photo-share-api/registry/local"
)

func main() {
	e := resolver.NewHandler(registry.NewProvider(local.InjectorConfig))
	e.Logger.Fatal(e.Start(":8080"))
}
