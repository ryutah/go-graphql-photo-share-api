package repository

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
)

type Tag interface {
	ListByUserIDs(context.Context, []model.UserID) (model.TagSlice, error)
	ListByPhotoIDs(context.Context, []model.PhotoID) (model.TagSlice, error)
}
