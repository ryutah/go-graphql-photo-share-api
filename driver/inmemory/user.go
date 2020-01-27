package inmemory

import (
	"context"

	"fmt"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
)

var userStorage = map[model.UserID]*model.User{
	"mHattrup": &model.User{
		ID:   "mHattrup",
		Name: "Mike Hattrup",
	},
	"gPlake": &model.User{
		ID:   "gPlake",
		Name: "Glen Plake",
	},
	"sSchmidt": &model.User{
		ID:   "sSchmidt",
		Name: "Scot Schmidt",
	},
}

type User struct{}

var (
	_ repository.User       = new(User)
	_ repository.UserSearch = new(User)
)

func NewUser() *User {
	return new(User)
}

func (u *User) Get(_ context.Context, id model.UserID) (*model.User, error) {
	user, ok := userStorage[id]
	if !ok {
		return nil, fmt.Errorf("not found user for id(%v)", id)
	}
	return user, nil
}

func (u *User) Search(ctx context.Context, q repository.UserQuery) ([]*model.User, error) {
	r := new(userQueryResolver)
	q.Resolve(r)

	results := make([]*model.User, 0, len(userStorage))
	for _, v := range userStorage {
		if val := r.inPhoto; val != nil && !existsTag(*val, v.ID) {
			continue
		}
		results = append(results, v)
	}

	return results, nil
}

type userQueryResolver struct {
	inPhoto *model.PhotoID
}

func (u *userQueryResolver) InPhoto(id model.PhotoID) {
	u.inPhoto = &id
}
