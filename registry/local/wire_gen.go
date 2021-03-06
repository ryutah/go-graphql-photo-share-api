// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package local

import (
	"github.com/google/wire"
	"github.com/ryutah/go-graphql-photo-share-api/domain/factory"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
	"github.com/ryutah/go-graphql-photo-share-api/driver/inmemory"
	"github.com/ryutah/go-graphql-photo-share-api/usecase"
)

// Injectors from local.go:

func InjectPhotoUsecase() *usecase.Photo {
	photo := inmemory.NewPhoto()
	factoryPhoto := factory.NewPhoto(photo)
	tag := inmemory.NewTag()
	usecasePhoto := usecase.NewPhoto(factoryPhoto, photo, photo, tag)
	return usecasePhoto
}

func InjectUserUsecase() *usecase.User {
	user := inmemory.NewUser()
	tag := inmemory.NewTag()
	usecaseUser := usecase.NewUser(user, user, tag)
	return usecaseUser
}

// local.go:

var dependencies = wire.NewSet(usecase.NewPhoto, usecase.NewUser, inmemory.NewPhoto, inmemory.NewUser, inmemory.NewTag, factory.NewPhoto, wire.Bind(new(repository.Photo), new(*inmemory.Photo)), wire.Bind(new(repository.PhotoSearch), new(*inmemory.Photo)), wire.Bind(new(repository.User), new(*inmemory.User)), wire.Bind(new(repository.UserSearch), new(*inmemory.User)), wire.Bind(new(repository.Tag), new(*inmemory.Tag)), wire.Bind(new(factory.PhotoIDGenerator), new(*inmemory.Photo)))
