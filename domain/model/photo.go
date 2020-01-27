package model

import (
	"fmt"
	"time"
)

type PhotoID string

type Photo struct {
	ID          PhotoID
	PostedBy    UserID
	Name        string
	Description string
	Category    PhotoCategory
	Created     time.Time
}

func (p *Photo) URL() URI {
	return URI(fmt.Sprintf("http://sample.com/%v.png", p.ID))
}
