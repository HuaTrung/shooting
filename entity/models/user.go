package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	ID         uint		`gorm:"primaryKey"`
	Name       string
	Email      string
	CreatedAt  *time.Time
	UpdatedAt  *time.Time
	DeletedAt  *time.Time
	stats  string
	Hashedpass string
	IsActive bool
}

func (user *User) GetId() uint {
	return user.ID
}

func (user *User) GetName() string {
	return user.Name
}