package model

import (
	"net/url"
)

type UserID string

type User struct {
	ID     UserID
	Name   string
	Avater *url.URL
}

func (u *User) AvaterURL() string {
	return u.Avater.String()
}
