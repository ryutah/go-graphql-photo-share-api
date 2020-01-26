package factory

import (
	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
)

type PhotoIDGenerator interface {
	NewID() model.PhotoID
}

type Photo struct {
	generator PhotoIDGenerator
}

func NewPhoto(g PhotoIDGenerator) *Photo {
	return &Photo{
		generator: g,
	}
}

func (p *Photo) FromPostPhoto(input model.PostPhotoInput) *model.Photo {
	photo := model.Photo{
		ID:   p.generator.NewID(),
		Name: input.Name,
	}
	if v := input.Category; v != nil {
		photo.Category = *v
	}
	if v := input.Description; v != nil {
		photo.Description = *v
	}
	return &photo
}
