package registry

import (
	"context"

	"github.com/ryutah/go-graphql-photo-share-api/usecase"
)

type InjectorConfig struct {
	Photo func() *usecase.Photo
	User  func() *usecase.User
}

type Provider struct {
	injectors InjectorConfig
}

func NewProvider(injectors InjectorConfig) *Provider {
	return &Provider{
		injectors: injectors,
	}
}

func (p *Provider) Photo(_ context.Context) *usecase.Photo {
	return p.injectors.Photo()
}

func (p *Provider) User(_ context.Context) *usecase.User {
	return p.injectors.User()
}
