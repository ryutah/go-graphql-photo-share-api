package resolver

import (
	"context"
	"errors"

	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql/dataloader"
	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

type photo struct {
	provier *registry.Provider
}

var _ graphql.PhotoResolver = new(photo)

func newPhoto(p *registry.Provider) *photo {
	return &photo{
		provier: p,
	}
}

func (p *photo) PostedBy(ctx context.Context, target *model.Photo) (*model.User, error) {
	return p.provier.User(ctx).Get(ctx, target.PostedBy)
}

func (p *photo) TaggedUsers(ctx context.Context, target *model.Photo) ([]*model.User, error) {
	loader, ok := dataloader.GetLoader(ctx)
	if !ok {
		return nil, errors.New("failed to get loader")
	}
	return loader.TaggedUser.Load(target.ID)
}
