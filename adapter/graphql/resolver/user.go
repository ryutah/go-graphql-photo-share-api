package resolver

import (
	"context"

	"errors"

	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql/dataloader"
	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

type user struct {
	provider *registry.Provider
}

var _ graphql.UserResolver = new(user)

func newUser(p *registry.Provider) *user {
	return &user{
		provider: p,
	}
}

func (u *user) PostedPhotos(ctx context.Context, target *model.User) ([]*model.Photo, error) {
	return u.provider.Photo(ctx).SearchPostedBy(ctx, target.ID)
}

func (u *user) InPhotos(ctx context.Context, target *model.User) ([]*model.Photo, error) {
	loader, ok := dataloader.GetLoader(ctx)
	if !ok {
		return nil, errors.New("failed to get loader")
	}
	return loader.InPhoto.Load(target.ID)
}
