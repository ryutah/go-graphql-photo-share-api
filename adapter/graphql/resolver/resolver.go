package resolver

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
)

type Root struct {
	query *query
}

var _ graphql.ResolverRoot = new(Root)

func NewRoot() *Root {
	return &Root{
		query: newQuery(),
	}
}

func (r *Root) Query() graphql.QueryResolver {
	return r.query
}

type query struct {
}

var _ graphql.QueryResolver = new(query)

func newQuery() *query {
	return new(query)
}

func (q *query) TotalPhotos(ctx context.Context) (int, error) {
	return 42, nil
}
