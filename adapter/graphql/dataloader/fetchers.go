package dataloader

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

func fetchPostedPhotos(ctx context.Context, p *registry.Provider) func([]model.UserID) ([][]*model.Photo, []error) {
	return func(ids []model.UserID) ([][]*model.Photo, []error) {
		userPhotos, err := p.Photo(ctx).SearchPostedBys(ctx, ids)
		if err != nil {
			return nil, []error{err}
		}
		results := make([][]*model.Photo, len(ids))
		for i, id := range ids {
			results[i] = userPhotos[id]
		}
		return results, nil
	}
}
func fetchPostedBy(ctx context.Context, p *registry.Provider) func([]model.UserID) ([]*model.User, []error) {
	return func(ids []model.UserID) ([]*model.User, []error) {
		users, err := p.User(ctx).List(ctx, ids)
		if err != nil {
			return nil, []error{err}
		}
		results := make([]*model.User, len(ids))
		for i, id := range ids {
			results[i] = users.Get(id)
		}
		return results, nil
	}
}

func fetchInPhoto(ctx context.Context, p *registry.Provider) func([]model.UserID) ([][]*model.Photo, []error) {
	return func(ids []model.UserID) ([][]*model.Photo, []error) {
		userPhotos, err := p.Photo(ctx).TaggedAsUsers(ctx, ids)
		if err != nil {
			return nil, []error{err}
		}

		results := make([][]*model.Photo, len(ids))
		for i, id := range ids {
			results[i] = userPhotos[id]
		}
		return results, nil
	}
}

func fetchTaggedUsers(ctx context.Context, p *registry.Provider) func([]model.PhotoID) ([][]*model.User, []error) {
	return func(ids []model.PhotoID) ([][]*model.User, []error) {
		photoUsers, err := p.User(ctx).InPhotos(ctx, ids)
		if err != nil {
			return nil, []error{err}
		}

		results := make([][]*model.User, len(ids))
		for i, id := range ids {
			results[i] = photoUsers[id]
		}
		return results, nil
	}
}
