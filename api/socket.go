package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
	"log"
	"shootingplane/database/repositories"
	"shootingplane/play/socket"
	"strconv"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
func SocketServiceRoom(c *gin.Context) {
	room_id, _ :=strconv.ParseUint(c.Param("id"),10, 32)
	var user_name string
	user_name, ok := c.GetQuery("user_name")
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if !ok {
		log.Print("Missing user")
		conn.WriteMessage(websocket.CloseMessage, []byte("Missing user_name parameter"))
		return
	}
	log.Printf("Someone's gonna loging into %s",room_id)
	db := c.MustGet("postgres").(*gorm.DB)
	wsserver:=c.MustGet("wsserver").(*socket.Platform)
	if err != nil {
		log.Println(err)
		conn.WriteMessage(websocket.CloseMessage, []byte("err"))
		return
	}
	user := repositories.FindUserByName(db,user_name)
	if user==nil{
		conn.WriteMessage(websocket.CloseMessage, []byte("your user_name's value is not existed"))
		return
	} else {
		client := socket.NewClient(conn, wsserver, user, uint(room_id))
		if client == nil {
			log.Fatal("ERROR WHEN LOGGING")
			conn.WriteMessage(websocket.CloseMessage, []byte("ERROR WHEN LOGGING"))
			return
		} else {
			room:=client.JoinRoom(uint(room_id), *user)
			if room==nil {
				conn.WriteMessage(websocket.CloseMessage, []byte("Room is not existed"))
				return
			} else {
				repositories.AddUserToRoom(db, uint(room_id),user.ID)

				wsserver.Register <- client
				go client.WritePump()
				go client.ReadPump()
			}
		}
	}
}

func SocketServicePlatform(c *gin.Context) {
	room_id, _ :=strconv.ParseUint(c.Param("id"),10, 32)
	var user_name string
	user_name, ok := c.GetQuery("user_name")
	if !ok {
		log.Fatal("Missing user")
		c.JSON(400,"Missing user")
	}
	log.Printf("Someone's gonna loging into %s",room_id)
	db := c.MustGet("postgres").(*gorm.DB)
	wsserver:=c.MustGet("wsserver").(*socket.Platform)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	user := repositories.FindUserByName(db,user_name)
	client := socket.NewClient(conn,wsserver ,user, uint(room_id))
	if client == nil{
		log.Fatal("ERROR WHEN LOGGING")
	} else {
		wsserver.Register <- client
		go client.WritePump()
		go client.ReadPump()
	}
}