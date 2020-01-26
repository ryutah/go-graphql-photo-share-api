package resolver

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

type mutation struct {
	provider *registry.Provider
}

var _ graphql.MutationResolver = new(mutation)

func newMutation(p *registry.Provider) *mutation {
	return &mutation{
		provider: p,
	}
}

func (m *mutation) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	return m.provider.Photo(ctx).Post(ctx, input)
}
