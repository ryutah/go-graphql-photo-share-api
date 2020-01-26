package model

import (
	"net/url"
)

type PhotoID string

type Photo struct {
	ID          PhotoID
	URL         *url.URL
	Name        string
	Description string
}

func (p *Photo) FullURL() string {
	if p.URL == nil {
		return ""
	}
	return p.URL.String()
}
