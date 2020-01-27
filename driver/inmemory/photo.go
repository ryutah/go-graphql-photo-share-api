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
	"github.com/ryutah/go-graphql-photo-share-api/lib/times"
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
			Created:     times.Date(1977, 3, 28, 0, 0, 0, 0),
		},
		"2": &model.Photo{
			ID:       "2",
			PostedBy: "sSchmidt",
			Name:     "Enjoying the sunshine",
			Category: model.PhotoCategorySelfie,
			Created:  times.Date(1985, 2, 1, 0, 0, 0, 0),
		},
		"3": &model.Photo{
			ID:          "3",
			PostedBy:    "sSchmidt",
			Description: "25 laps on gunbarrel today",
			Name:        "Gunbarrel 25",
			Created:     times.Date(2018, 4, 15, 19, 9, 57, 0),
		},
	}
)

type Photo struct{}

var (
	_ repository.Photo         = new(Photo)
	_ repository.PhotoSearch   = new(Photo)
	_ factory.PhotoIDGenerator = new(Photo)
)

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

func (p *Photo) Search(ctx context.Context, q repository.PhotoQuery) ([]*model.Photo, error) {
	r := new(photoQueryResolver)
	q.Reslove(r)

	result := make([]*model.Photo, 0, len(photoStorage))
	for _, v := range photoStorage {
		if val := r.postedBy; val != nil && v.PostedBy != *val {
			continue
		}
		if val := r.tagged; val != nil && !existsTag(v.ID, *val) {
			continue
		}
		result = append(result, v)
	}

	return result, nil
}

func (p *Photo) NewID() model.PhotoID {
	photoStorageMux.Lock()
	defer photoStorageMux.Unlock()
	return model.PhotoID(strconv.Itoa(len(photoStorage) + 1))
}

type photoQueryResolver struct {
	postedBy *model.UserID
	tagged   *model.UserID
}

func (p *photoQueryResolver) PostedBy(id model.UserID) {
	p.postedBy = &id
}

func (p *photoQueryResolver) Tagged(id model.UserID) {
	p.tagged = &id
}
