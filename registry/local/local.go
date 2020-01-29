//+build wireinject

package local

import (
	"github.com/google/wire"
	"github.com/ryutah/go-graphql-photo-share-api/domain/factory"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
	"github.com/ryutah/go-graphql-photo-share-api/driver/inmemory"
	"github.com/ryutah/go-graphql-photo-share-api/usecase"
)

var dependencies = wire.NewSet(
	usecase.NewPhoto,
	usecase.NewUser,
	inmemory.NewPhoto,
	inmemory.NewUser,
	inmemory.NewTag,
	factory.NewPhoto,
	wire.Bind(new(repository.Photo), new(*inmemory.Photo)),
	wire.Bind(new(repository.PhotoSearch), new(*inmemory.Photo)),
	wire.Bind(new(repository.User), new(*inmemory.User)),
	wire.Bind(new(repository.UserSearch), new(*inmemory.User)),
	wire.Bind(new(repository.Tag), new(*inmemory.Tag)),
	wire.Bind(new(factory.PhotoIDGenerator), new(*inmemory.Photo)),
)

func InjectPhotoUsecase() *usecase.Photo {
	panic(wire.Build(dependencies))
}

func InjectUserUsecase() *usecase.User {
	panic(wire.Build(dependencies))
}
