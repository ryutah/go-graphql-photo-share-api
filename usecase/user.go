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
	}
}

func NewUser(userRepo repository.User, userSearch repository.UserSearch) *User {
	return &User{
		repository: struct {
			user       repository.User
			userSearch repository.UserSearch
		}{
			user:       userRepo,
			userSearch: userSearch,
		},
	}
}

func (u *User) Get(ctx context.Context, id model.UserID) (*model.User, error) {
	return u.repository.user.Get(ctx, id)
}

func (u *User) InPhoto(ctx context.Context, id model.PhotoID) ([]*model.User, error) {
	return u.repository.userSearch.Search(
		ctx, repository.CreateUserQuery().WithInPhoto(id),
	)
}
