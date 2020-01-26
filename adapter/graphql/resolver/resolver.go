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
	photo    *photo
	user     *user
}

var _ graphql.ResolverRoot = new(Root)

func NewRoot() *Root {
	return &Root{
		query:    newQuery(),
		mutation: newMutation(),
		photo:    newPhoto(),
		user:     newUser(),
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
				PostedBy:    "user1",
			},
			{
				ID:          "photo2",
				URL:         mustURL("http://sample.com/photo2.png"),
				Name:        "Photo2",
				Description: "description of photo2",
				Category:    model.PhotoCategoryLandscape,
				PostedBy:    "user2",
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

type photo struct{}

var _ graphql.PhotoResolver = new(photo)

func newPhoto() *photo {
	return new(photo)
}

func (p *photo) PostedBy(ctx context.Context, target *model.Photo) (*model.User, error) {
	return &model.User{
		ID:     target.PostedBy,
		Name:   "User",
		Avater: mustURL("http://sample.com/avater.png"),
	}, nil
}

type user struct{}

var _ graphql.UserResolver = new(user)

func newUser() *user {
	return new(user)
}

func (u *user) PostedPhotos(ctx context.Context, target *model.User) ([]*model.Photo, error) {
	return []*model.Photo{
		{
			ID:          "photo1",
			PostedBy:    target.ID,
			URL:         mustURL("http://sample.com/photo1.png"),
			Name:        "photo1",
			Description: "description",
			Category:    model.PhotoCategoryLandscape,
		},
		{
			ID:          "photo2",
			PostedBy:    target.ID,
			URL:         mustURL("http://sample.com/photo2.png"),
			Name:        "photo2",
			Description: "description",
			Category:    model.PhotoCategoryGraphic,
		},
	}, nil
}

func mustURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
