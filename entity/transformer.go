package entity

import (
	"shootingplane/entity/api"
	"shootingplane/entity/models"
)

func ConvertToRoomAPI(r *models.Room) *api.Room {
	return &api.Room{
		ID: r.ID,
		Name: r.Name,
		Password: r.Password,
	}
}

func ConvertToRoomModel(r *api.Room) *models.Room {
	return &models.Room{
		Name: r.Name,
		Password: r.Password,
	}
}