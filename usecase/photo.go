package usecase

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/factory"
	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
)

type Photo struct {
	factory struct {
		photo *factory.Photo
	}
	repository struct {
		photo repository.Photo
	}
}

func NewPhoto(photoFact *factory.Photo, photoRepo repository.Photo) *Photo {
	return &Photo{
		factory: struct {
			photo *factory.Photo
		}{
			photo: photoFact,
		},
		repository: struct {
			photo repository.Photo
		}{
			photo: photoRepo,
		},
	}
}

func (p *Photo) Post(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	newPhoto := p.factory.photo.FromPostPhoto(input)
	if err := p.repository.photo.Create(ctx, *newPhoto); err != nil {
		return nil, err
	}
	return newPhoto, nil
}

func (p *Photo) All(ctx context.Context) ([]*model.Photo, error) {
	return p.repository.photo.All(ctx)
}

func (p *Photo) TotalCount(ctx context.Context) (int, error) {
	return p.repository.photo.Count(ctx)
}

func (p *Photo) SearchPostedBy(ctx context.Context, postedBy model.UserID) ([]*model.Photo, error) {
	return p.repository.photo.ListPostedBy(ctx, postedBy)
}
