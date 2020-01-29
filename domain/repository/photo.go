package repository

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/domain/model"
)

type Photo interface {
	Create(context.Context, model.Photo) error
	All(context.Context) (model.PhotoList, error)
	Count(context.Context) (int, error)
	GetMulti(context.Context, []model.PhotoID) (model.PhotoList, error)
}

type PhotoSearch interface {
	Search(context.Context, PhotoQuery) (model.PhotoList, error)
}

type PhotoQueryResolver interface {
	PostedBys(...model.UserID)
	Tagged(model.UserID)
}

type PhotoQuery struct {
	postedBys []model.UserID
	tagged    *model.UserID
}

func CreatePhotoQuery() PhotoQuery {
	return PhotoQuery{}
}

func (p PhotoQuery) WithPostedBys(id ...model.UserID) PhotoQuery {
	p.postedBys = append(p.postedBys, id...)
	return p
}

func (p PhotoQuery) WithTagged(id model.UserID) PhotoQuery {
	p.tagged = &id
	return p
}

func (p PhotoQuery) Reslove(r PhotoQueryResolver) {
	if v := p.postedBys; v != nil {
		r.PostedBys(v...)
	}
	if v := p.tagged; v != nil {
		r.Tagged(*v)
	}
}
