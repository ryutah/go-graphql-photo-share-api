package model

import (
	"fmt"
	"sort"
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

type PhotoList map[PhotoID]*Photo

func (p PhotoList) Add(photo *Photo) {
	p[photo.ID] = photo
}

func (p PhotoList) Get(id PhotoID) *Photo {
	return p[id]
}

func (p PhotoList) Slice() []*Photo {
	results := make([]*Photo, 0, len(p))
	for _, photo := range p {
		results = append(results, photo)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].ID < results[j].ID
	})
	return results
}
