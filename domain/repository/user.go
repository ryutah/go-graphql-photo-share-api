package repository

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
)

type User interface {
	Get(context.Context, model.UserID) (*model.User, error)
	GetMulti(context.Context, []model.UserID) (model.UserList, error)
}

type UserSearch interface {
	Search(context.Context, UserQuery) (model.UserList, error)
}

type UserQueryResolver interface {
	InPhoto(model.PhotoID)
}

type UserQuery struct {
	inPhoto *model.PhotoID
}

func CreateUserQuery() UserQuery {
	return UserQuery{}
}

func (u UserQuery) WithInPhoto(id model.PhotoID) UserQuery {
	u.inPhoto = &id
	return u
}

func (u UserQuery) Resolve(r UserQueryResolver) {
	if v := u.inPhoto; v != nil {
		r.InPhoto(*v)
	}
}
