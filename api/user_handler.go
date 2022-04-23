package api

import (
	"gorm.io/gorm"
	"net/http"
	"shootingplane/database/repositories"
	"shootingplane/entity"
	"shootingplane/entity/api"
	"shootingplane/play/socket"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// CreateUser : Controller for creating user
func CreateUser(c *gin.Context) {
	var param api.User
	if err := c.BindJSON(&param); !errors.Is(err, nil) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		//db := c.MustGet("postgres").(*gorm.DB)
		//user := &models.User{
		//	Name:  param.Name,
		//	Email: param.Email,
		//	IsActive:true,
		//}
		//result := repositories.AddUser(db,user)
		//if result.Error != nil {
		//	panic(err)
		//}
		c.JSON(200, "user")
	}
}

// CreateRoom : Controller for creating user
func CreateRoom(c *gin.Context) {
	var param api.Room
	if err := c.BindJSON(&param); !errors.Is(err, nil) {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	} else {
		db := c.MustGet("postgres").(*gorm.DB)
		server := c.MustGet("wsserver").(*socket.Platform)

		res:=repositories.CreateRoom(db,entity.ConvertToRoomModel(&param))
		if res.ID == 0 {
			c.JSON(200, "No available room")
		} else {
			// generate socket for a room
			socket.NewRoom(server,res.Name,res.ID)
			c.JSON(200, res)
		}

	}
}
