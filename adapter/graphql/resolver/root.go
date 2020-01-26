package resolver

import (
	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

type Root struct {
	query    *query
	mutation *mutation
	photo    *photo
	user     *user
}

var _ graphql.ResolverRoot = new(Root)

func NewRoot(p *registry.Provider) *Root {
	return &Root{
		query:    newQuery(p),
		mutation: newMutation(p),
		photo:    newPhoto(p),
		user:     newUser(p),
	}
}

func (r *Root) Query() graphql.QueryResolver {
	return r.query
}

func (r *Root) Mutation() graphql.MutationResolver {
	return r.mutation
}

func (r *Root) Photo() graphql.PhotoResolver {
	return r.photo
}

func (r *Root) User() graphql.UserResolver {
	return r.user
}
