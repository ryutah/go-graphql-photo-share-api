package repository

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
)

type User interface {
	Get(context.Context, model.UserID) (*model.User, error)
}
