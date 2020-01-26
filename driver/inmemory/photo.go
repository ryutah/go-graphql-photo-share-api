package inmemory

import (
	"context"
	"sync"

	"fmt"

	"strconv"

	"sort"

	"github.com/ryutah/go-graphql-photo-share-api/domain/factory"
	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
	"github.com/ryutah/go-graphql-photo-share-api/domain/repository"
)

var (
	photoStorageMux = new(sync.Mutex)
	photoStorage    = map[model.PhotoID]*model.Photo{
		"1": &model.Photo{
			ID:          "1",
			PostedBy:    "gPlake",
			Name:        "Dropping the Heart Chute",
			Description: "The heart chute is one of my favorite chutes",
			Category:    model.PhotoCategoryAction,
		},
		"2": &model.Photo{
			ID:       "2",
			PostedBy: "sSchmidt",
			Name:     "Enjoying the sunshine",
			Category: model.PhotoCategorySelfie,
		},
		"3": &model.Photo{
			ID:          "3",
			PostedBy:    "sSchmidt",
			Description: "25 laps on gunbarrel today",
			Name:        "Gunbarrel 25",
			Category:    model.PhotoCategoryLandscape,
		},
	}
)

type Photo struct{}

var _ repository.Photo = new(Photo)
var _ factory.PhotoIDGenerator = new(Photo)

func NewPhoto() *Photo {
	return new(Photo)
}

func (p *Photo) Create(ctx context.Context, photo model.Photo) error {
	photoStorageMux.Lock()
	defer photoStorageMux.Unlock()

	if _, ok := photoStorage[photo.ID]; ok {
		return fmt.Errorf("deprecate photo id(%v)", photo.ID)
	}
	photoStorage[photo.ID] = &photo
	return nil
}

func (p *Photo) All(_ context.Context) ([]*model.Photo, error) {
	result := make([]*model.Photo, 0, len(photoStorage))
	for _, v := range photoStorage {
		result = append(result, v)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].ID < result[j].ID
	})
	return result, nil
}

func (p *Photo) Count(_ context.Context) (int, error) {
	return len(photoStorage), nil
}

func (p *Photo) ListPostedBy(_ context.Context, postedBy model.UserID) ([]*model.Photo, error) {
	result := make([]*model.Photo, 0, len(photoStorage))
	for _, v := range photoStorage {
		if v.PostedBy == postedBy {
			result = append(result, v)
		}
	}
	return result, nil
}

func (p *Photo) NewID() model.PhotoID {
	photoStorageMux.Lock()
	defer photoStorageMux.Unlock()
	return model.PhotoID(strconv.Itoa(len(photoStorage) + 1))
}
