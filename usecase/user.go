package usecase

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
)

type User struct {
	repository struct {
		user       repository.User
		userSearch repository.UserSearch
		tag        repository.Tag
	}
}

func NewUser(userRepo repository.User, userSearch repository.UserSearch, tagRepo repository.Tag) *User {
	return &User{
		repository: struct {
			user       repository.User
			userSearch repository.UserSearch
			tag        repository.Tag
		}{
			user:       userRepo,
			userSearch: userSearch,
			tag:        tagRepo,
		},
	}
}

func (u *User) Get(ctx context.Context, id model.UserID) (*model.User, error) {
	return u.repository.user.Get(ctx, id)
}

func (u *User) List(ctx context.Context, ids []model.UserID) (model.UserList, error) {
	return u.repository.user.GetMulti(ctx, ids)
}

func (u *User) InPhotos(ctx context.Context, ids []model.PhotoID) (map[model.PhotoID][]*model.User, error) {
	tags, err := u.repository.tag.ListByPhotoIDs(ctx, ids)
	if err != nil {
		return nil, err
	}
	users, err := u.repository.user.GetMulti(ctx, tags.UserIDs())
	if err != nil {
		return nil, err
	}
	results := make(map[model.PhotoID][]*model.User)
	for _, t := range tags {
		results[t.PhotoID] = append(results[t.PhotoID], users.Get(t.UserID))
	}
	return results, nil
}
