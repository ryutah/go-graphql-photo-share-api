package model

type UserID string

type User struct {
	ID     UserID
	Name   string
	Avatar URI
}

type UserList map[UserID]*User

func (u UserList) Add(user *User) {
	u[user.ID] = user
}

func (u UserList) Get(id UserID) *User {
	return u[id]
}

func (u UserList) Slice() []*User {
	results := make([]*User, 0, len(u))
	for _, user := range u {
		results = append(results, user)
	}
	return results
}
