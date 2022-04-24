package repositories

import (
	"gorm.io/gorm"
	"log"
	"shootingplane/entity"
	"shootingplane/entity/api"
	"shootingplane/entity/models"
)

func CreateRoom(Db *gorm.DB,room *models.Room) *api.Room{
	tb := Db.Table("personal.rooms")
	var new_room models.Room
	// Query at least one free room.
	res:=tb.Unscoped().First(&new_room,"\"is_acquired\" = ?", "false")
	if res.Error!= nil {
		log.Fatal(res)
	} else {
		new_room.Name=room.Name
		if room.HasPassword == true {
			new_room.Password=room.Password
		}
		res=tb.Unscoped().Save(&new_room)
	}
	return entity.ConvertToRoomAPI(&new_room)
}

func AddUserToRoom(Db *gorm.DB,room_id uint,client_id uint){
	tb := Db.Table("personal.rooms")
	var new_room models.Room
	// Query at least one free room.
	res:=tb.Unscoped().First(&new_room,"\"id\" = ?", room_id)
	if res.Error!= nil {
		log.Fatal(res)
	} else {
		if new_room.PlayerA==0{
			new_room.PlayerA=client_id
			res=tb.Unscoped().Save(&new_room)
			return
		}
		if new_room.PlayerB==0{
			new_room.PlayerB=client_id
			res=tb.Unscoped().Save(&new_room)
			return
		}
	}
}
func 	RemoveUserFromRoom(Db *gorm.DB,room_id uint,client_id uint){
	tb := Db.Table("personal.rooms")
	var new_room models.Room
	// Query at least one free room.
	res:=tb.Unscoped().First(&new_room,"\"id\" = ?", room_id)
	if res.Error!= nil {
		log.Fatal(res)
	} else {
		if new_room.PlayerA==client_id{
			new_room.PlayerA=0
			res=tb.Unscoped().Save(&new_room)
			return
		}
		if new_room.PlayerB==client_id{
			new_room.PlayerB=0
			res=tb.Unscoped().Save(&new_room)
			return
		}
	}
}