package local

import (
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

var InjectorConfig = registry.InjectorConfig{
	Photo: InjectPhotoUsecase,
	User:  InjectUserUsecase,
}
