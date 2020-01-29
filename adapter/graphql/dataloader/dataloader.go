//go:generate go run github.com/vektah/dataloaden TaggedUsersLoader github.com/ryutah/go-graphql-photo-share-api/domain/model.PhotoID []*github.com/ryutah/go-graphql-photo-share-api/domain/model.User
//go:generate go run github.com/vektah/dataloaden InPhotoLoader github.com/ryutah/go-graphql-photo-share-api/domain/model.UserID []*github.com/ryutah/go-graphql-photo-share-api/domain/model.Photo

package dataloader

import (
	"time"

	"context"

	"github.com/labstack/echo/v4"
	"github.com/ryutah/go-graphql-photo-share-api/registry"
)

const (
	wait     = 50 * time.Millisecond
	maxBatch = 100
)

type loaderKey struct{}

func GetLoader(ctx context.Context) (*Loaders, bool) {
	l, ok := ctx.Value(loaderKey{}).(*Loaders)
	return l, ok
}

type Loaders struct {
	InPhoto    *InPhotoLoader
	TaggedUser *TaggedUsersLoader
}

func Middleware(p *registry.Provider) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			ctx := c.Request().Context()

			l := &Loaders{
				InPhoto: NewInPhotoLoader(InPhotoLoaderConfig{
					Fetch:    fetchInPhoto(ctx, p),
					Wait:     wait,
					MaxBatch: maxBatch,
				}),
				TaggedUser: NewTaggedUsersLoader(TaggedUsersLoaderConfig{
					Fetch:    fetchTaggedUsers(ctx, p),
					Wait:     wait,
					MaxBatch: maxBatch,
				}),
			}

			newCtx := context.WithValue(ctx, loaderKey{}, l)
			c.SetRequest(c.Request().WithContext(newCtx))

			return next(c)
		}
	}
}
