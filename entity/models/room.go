package models

import (
	"gorm.io/gorm"
	"shootingplane/entity/api"
	"time"
)

type Room struct {
	gorm.Model
	ID         uint		`gorm:"primaryKey"`
	Name       string
	Game       string
	HasPassword bool
	IsAcquired bool
	CreatedAt  *time.Time
	Password 	string
}

func (r *Room) ConvertToRoomAPI() *api.Room {
	return &api.Room{
		ID: r.ID,
		Name: r.Name,
		Password: r.Password,
	}
}