package usecase

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
)

type User struct {
	repository struct {
		user repository.User
	}
}

func NewUser(userRepo repository.User) *User {
	return &User{
		repository: struct {
			user repository.User
		}{
			user: userRepo,
		},
	}
}

func (u *User) Get(ctx context.Context, id model.UserID) (*model.User, error) {
	return u.repository.user.Get(ctx, id)
}
