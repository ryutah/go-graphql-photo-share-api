package resolver

import (
	"context"
	"net/url"

	"github.com/ryutah/go-graphql-photo-share-api/adapter/graphql"
	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
)

type Root struct {
	query    *query
	mutation *mutation
}

var _ graphql.ResolverRoot = new(Root)

func NewRoot() *Root {
	return &Root{
		query:    newQuery(),
		mutation: newMutation(),
	}
}

func (r *Root) Query() graphql.QueryResolver {
	return r.query
}

func (r *Root) Mutation() graphql.MutationResolver {
	return r.mutation
}

type query struct {
	photos []*model.Photo
}

var _ graphql.QueryResolver = new(query)

func newQuery() *query {
	return &query{
		photos: []*model.Photo{
			{
				ID:          "photo1",
				URL:         mustURL("http://sample.com/photo1.png"),
				Name:        "Photo1",
				Description: "description of photo1",
				Category:    model.PhotoCategorySelfie,
			},
			{
				ID:          "photo2",
				URL:         mustURL("http://sample.com/photo2.png"),
				Name:        "Photo2",
				Description: "description of photo2",
				Category:    model.PhotoCategoryLandscape,
			},
		},
	}
}

func (q *query) TotalPhotos(ctx context.Context) (int, error) {
	return len(q.photos), nil
}

func (q *query) AllPhotos(ctx context.Context) ([]*model.Photo, error) {
	return q.photos, nil
}

type mutation struct {
}

var _ graphql.MutationResolver = new(mutation)

func newMutation() *mutation {
	return new(mutation)
}

func (m *mutation) PostPhoto(ctx context.Context, input model.PostPhotoInput) (*model.Photo, error) {
	newPhoto := &model.Photo{
		ID:       "new_photo",
		Name:     input.Name,
		Category: *input.Category,
		URL:      mustURL("http://sample.com/new_photo.png"),
	}
	if input.Description != nil {
		newPhoto.Description = *input.Description
	}
	return newPhoto, nil
}

func mustURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
