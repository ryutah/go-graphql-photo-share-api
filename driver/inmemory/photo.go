package inmemory

import (
	"context"
	"sync"
	"time"

	"fmt"

	"strconv"

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

func (p *Photo) All(_ context.Context) (model.PhotoList, error) {
	result := make(model.PhotoList)
	for _, v := range photoStorage {
		result.Add(v)
	}
	return result, nil
}

func (p *Photo) Count(_ context.Context) (int, error) {
	return len(photoStorage), nil
}

func (p *Photo) GetMulti(_ context.Context, ids []model.PhotoID) (model.PhotoList, error) {
	results := make(model.PhotoList)
	for _, id := range ids {
		if photo, ok := photoStorage[id]; ok {
			results.Add(photo)
		}
	}
	return results, nil
}

func (p *Photo) Search(ctx context.Context, q repository.PhotoQuery) (model.PhotoList, error) {
	r := new(photoQueryResolver)
	q.Reslove(r)

	result := make(model.PhotoList)
	for _, v := range photoStorage {
		if r.isMatchQuery(v) {
			result.Add(v)
		}
	}

	return result, nil
}

func (p *Photo) NewID() model.PhotoID {
	photoStorageMux.Lock()
	defer photoStorageMux.Unlock()
	return model.PhotoID(strconv.Itoa(len(photoStorage) + 1))
}

type photoQueryResolver struct {
	postedBys    []model.UserID
	tagged       *model.UserID
	createdAfter *time.Time
}

func (p *photoQueryResolver) PostedBys(id ...model.UserID) {
	p.postedBys = append(p.postedBys, id...)
}

func (p *photoQueryResolver) Tagged(id model.UserID) {
	p.tagged = &id
}

func (p *photoQueryResolver) CreatedAfter(t time.Time) {
	p.createdAfter = &t
}

func (p *photoQueryResolver) isMatchQuery(photo *model.Photo) bool {
	if photo == nil {
		return false
	}

	var (
		isMatchPostedBy = func(postedBy model.UserID) bool {
			if len(p.postedBys) == 0 {
				return true
			}
			for _, id := range p.postedBys {
				if postedBy == id {
					return true
				}
			}
			return false
		}
		isMatchTaggedUser = func(id model.PhotoID) bool {
			if p.tagged == nil {
				return true
			}
			return tagStorage.exists(id, *p.tagged)
		}
		isMatchCreatedAfter = func(t time.Time) bool {
			if p.createdAfter == nil {
				return true
			}
			return t.Equal(*p.createdAfter) || t.After(*p.createdAfter)
		}
	)

	return isMatchPostedBy(photo.PostedBy) && isMatchTaggedUser(photo.ID) && isMatchCreatedAfter(photo.Created)
}
