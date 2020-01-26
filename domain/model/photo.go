package model

import (
	"fmt"
)

type PhotoID string

type Photo struct {
	ID          PhotoID
	PostedBy    UserID
	Name        string
	Description string
	Category    PhotoCategory
}

func (p *Photo) URL() URI {
	return URI(fmt.Sprintf("http://sample.com/%v.png", p.ID))
}
