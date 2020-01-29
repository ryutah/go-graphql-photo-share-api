package resolver

import (
	"context"

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
	return dataloader.GetLoader(ctx).PostedBy.Load(target.PostedBy)
}

func (p *photo) TaggedUsers(ctx context.Context, target *model.Photo) ([]*model.User, error) {
	return dataloader.GetLoader(ctx).TaggedUser.Load(target.ID)
}
