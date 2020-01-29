package resolver

import (
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

type root struct {
	query    *query
	mutation *mutation
	photo    *photo
	user     *user
}

var _ graphql.ResolverRoot = new(root)

func newRoot(p *registry.Provider) *root {
	return &root{
		query:    newQuery(p),
		mutation: newMutation(p),
		photo:    newPhoto(p),
		user:     newUser(p),
	}
}

func (r *root) Query() graphql.QueryResolver {
	return r.query
}

func (r *root) Mutation() graphql.MutationResolver {
	return r.mutation
}

func (r *root) Photo() graphql.PhotoResolver {
	return r.photo
}

func (r *root) User() graphql.UserResolver {
	return r.user
}
