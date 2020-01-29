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
		tag         repository.Tag
	}
}

func NewPhoto(photoFact *factory.Photo, photoRepo repository.Photo, photoSearch repository.PhotoSearch, tagRepo repository.Tag) *Photo {
	return &Photo{
		factory: struct {
			photo *factory.Photo
		}{
			photo: photoFact,
		},
		repository: struct {
			photo       repository.Photo
			photoSearch repository.PhotoSearch
			tag         repository.Tag
		}{
			photo:       photoRepo,
			photoSearch: photoSearch,
			tag:         tagRepo,
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
	photos, err := p.repository.photo.All(ctx)
	if err != nil {
		return nil, err
	}
	return photos.Slice(), nil
}

func (p *Photo) TotalCount(ctx context.Context) (int, error) {
	return p.repository.photo.Count(ctx)
}

func (p *Photo) SearchPostedBys(ctx context.Context, postedBys []model.UserID) (map[model.UserID][]*model.Photo, error) {
	photos, err := p.repository.photoSearch.Search(
		ctx, repository.CreatePhotoQuery().WithPostedBys(postedBys...),
	)
	if err != nil {
		return nil, err
	}
	results := make(map[model.UserID][]*model.Photo)
	for _, photo := range photos {
		results[photo.PostedBy] = append(results[photo.PostedBy], photo)
	}
	return results, nil
}

func (p *Photo) TaggedAsUsers(ctx context.Context, ids []model.UserID) (map[model.UserID][]*model.Photo, error) {
	tags, err := p.repository.tag.ListByUserIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	photos, err := p.repository.photo.GetMulti(ctx, tags.PhotoIDs())
	if err != nil {
		return nil, err
	}
	results := make(map[model.UserID][]*model.Photo)
	for _, t := range tags {
		results[t.UserID] = append(results[t.UserID], photos.Get(t.PhotoID))
	}
	return results, nil
}
