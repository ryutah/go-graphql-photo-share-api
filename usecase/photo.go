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
		photo       repository.Photo
		photoSearch repository.PhotoSearch
	}
}

func NewPhoto(photoFact *factory.Photo, photoRepo repository.Photo, photoSearch repository.PhotoSearch) *Photo {
	return &Photo{
		factory: struct {
			photo *factory.Photo
		}{
			photo: photoFact,
		},
		repository: struct {
			photo       repository.Photo
			photoSearch repository.PhotoSearch
		}{
			photo:       photoRepo,
			photoSearch: photoSearch,
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
	return p.repository.photoSearch.Search(
		ctx, repository.CreatePhotoQuery().WithPostedBy(postedBy),
	)
}

func (p *Photo) Tagged(ctx context.Context, taggedUser model.UserID) ([]*model.Photo, error) {
	return p.repository.photoSearch.Search(
		ctx, repository.CreatePhotoQuery().WithTagged(taggedUser),
	)
}
