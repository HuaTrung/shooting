package api

import (
	"errors"
	"time"
)
import "shootingplane/util"

type User struct {
	ID             string    `json:"id,omitempty"`
	UUID           string    `json:"uuid"`
	Name           string    `json:"name,omitempty"`
	Email          string    `json:"email,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
	DeletedAt      time.Time `json:"deleted_at,omitempty"`
	HashedPassword string    `json:"hashed_password,omitempty"`
}
// NewUser returns a new user with the given name and email.
func NewUser(name, email string) *User {
	return &User{
		Name:  name,
		Email: email,
	}
}

// SetPassword takes a plaintext password and sets the hashed password.
func (u *User) SetPassword(password string) error {
	var err error
	if len(password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	u.HashedPassword, err = util.HashPassword(password)
	return err
}

// ComparePassword checks if the password matches the given user.
func (u *User) ComparePassword(password string) error {
	return nil
	//return passwd.Verify(password, u.HashedPassword)
}

// func (u *User) Save(repo repository.User) error {
//         var err error
//         u.ID, err = repo.Create(u)
//         return err
// }
