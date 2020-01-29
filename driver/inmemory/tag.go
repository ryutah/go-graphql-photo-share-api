package inmemory

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
)

var tagStorage = tagSlice{
	{PhotoID: "1", UserID: "gPlake"},
	{PhotoID: "2", UserID: "sSchmidt"},
	{PhotoID: "2", UserID: "mHattrup"},
	{PhotoID: "2", UserID: "gPlake"},
}

type Tag struct{}

var _ repository.Tag = new(Tag)

func NewTag() *Tag {
	return new(Tag)
}

func (t *Tag) ListByUserIDs(ctx context.Context, ids []model.UserID) (model.TagSlice, error) {
	var results model.TagSlice
	for _, id := range ids {
		tags := tagStorage.byUserID(id)
		results = append(results, tags...)
	}
	return results, nil
}

func (t *Tag) ListByPhotoIDs(ctx context.Context, ids []model.PhotoID) (model.TagSlice, error) {
	var results model.TagSlice
	for _, id := range ids {
		tags := tagStorage.byPhotoID(id)
		results = append(results, tags...)
	}
	return results, nil
}

type tagSlice []*model.Tag

func (t tagSlice) exists(photoID model.PhotoID, userID model.UserID) bool {
	extracted := t.byPhotoID(photoID)
	if len(extracted) == 0 {
		return false
	}
	return len(extracted.byUserID(userID)) != 0
}

func (t tagSlice) byPhotoID(id model.PhotoID) tagSlice {
	var results tagSlice
	for _, tg := range t {
		if tg.PhotoID == id {
			results = append(results, tg)
		}
	}
	return results
}

func (t tagSlice) byUserID(id model.UserID) tagSlice {
	var results tagSlice
	for _, tg := range t {
		if tg.UserID == id {
			results = append(results, tg)
		}
	}
	return results
}
