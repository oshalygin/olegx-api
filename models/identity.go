package models

// Identity represents the user id
type Identity interface {
	GetID() string
	GetName() string
}

// User represents the logged in user
type User struct {
	ID   string
	Name string
}

// GetID retrieves the ID of the user
func (u User) GetID() string {
	return u.ID
}

// GetName retrieves the name of the user
func (u User) GetName() string {
	return u.Name
}
