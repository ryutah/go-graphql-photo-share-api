package repository

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
)

type Photo interface {
	Create(context.Context, model.Photo) error
	All(context.Context) ([]*model.Photo, error)
	Count(context.Context) (int, error)
	ListPostedBy(context.Context, model.UserID) ([]*model.Photo, error)
}
