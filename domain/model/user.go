package model

type UserID string

type User struct {
	ID     UserID
	Name   string
	Avatar URI
}
