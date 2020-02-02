package resolver

import (
	"context"
	"time"

	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

type query struct {
	provider *registry.Provider
}

var _ graphql.QueryResolver = new(query)

func newQuery(p *registry.Provider) *query {
	return &query{
		provider: p,
	}
}

func (q *query) TotalPhotos(ctx context.Context) (int, error) {
	return q.provider.Photo(ctx).TotalCount(ctx)
}

func (q *query) AllPhotos(ctx context.Context, after *time.Time) ([]*model.Photo, error) {
	query := repository.CreatePhotoQuery()
	if after != nil {
		query = query.WithCreatedAfter(*after)
	}
	return q.provider.Photo(ctx).All(ctx, query)
}
