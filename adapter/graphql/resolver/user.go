package resolver

import (
	"context"

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
	return dataloader.GetLoader(ctx).PostedPhoto.Load(target.ID)
}

func (u *user) InPhotos(ctx context.Context, target *model.User) ([]*model.Photo, error) {
	return dataloader.GetLoader(ctx).InPhoto.Load(target.ID)
}
