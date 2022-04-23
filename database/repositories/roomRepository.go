package repositories

import (
	"gorm.io/gorm"
	"shootingplane/entity"
	"shootingplane/entity/api"
	"shootingplane/entity/models"
	"log"
)

func CreateRoom(Db *gorm.DB,room *models.Room) *api.Room{
	tb := Db.Table("personal.rooms")
	var new_room models.Room
	// Query at least one free room.
	res:=tb.Unscoped().First(&new_room,"\"IsAcquired\" = ?", "false")
	if res.Error!= nil {
		log.Fatal(res)
	} else {
		new_room.Name=room.Name
		if room.HasPassword == true {
			new_room.Password=room.Password
		}
		tb.Save(&new_room)
	}
	return entity.ConvertToRoomAPI(&new_room)
}
