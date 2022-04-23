package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"shootingplane/play/socket"
	"strconv"
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
func SocketService(c *gin.Context) {
	id, _ :=strconv.ParseUint(c.Param("id"),10, 32)
	log.Printf("Someone's gonna loging into %s",id)
	wsserver:=c.MustGet("wsserver").(*socket.Platform)
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := socket.NewClient(conn,wsserver ,"trung", uint(id))
	if client == nil{
		log.Fatal("ERROR WHEN LOGGING")
	} else {
		wsserver.Register <- client
		go client.WritePump()
		go client.ReadPump()
	}
}
