package main

import (
	"github.com/gin-gonic/gin"
	"shootingplane/api"
	"shootingplane/database/config"
	"shootingplane/play/socket"
)

func main() {
	pg:=database.GetPgClient()
	server := socket.NewPlatformServer(pg)
	go server.Run()
	// initialize new gin engine (for server)
	r := gin.Default()

	// configure database
	// mongo, _ := models.NewMongo("mongodb+srv://trung:6IEE17hn2qTyR66x@shooting.nlwel.mongodb.net/shooting?retryWrites=true&w=majority")

	// configure firebase
	// firebaseAuth := models.SetupFirebase()

	// set models & firebase auth to gin context with a middleware to all incoming request

	r.Use(func(c *gin.Context) {
		c.Set("postgres",pg)
		c.Set("wsserver",server)
	})

	/* ---------------------------  Public routes  --------------------------- */

	r.POST("/user", api.CreateUser)
	r.POST("/room", api.CreateRoom)
	r.GET("room/:id/ws", api.SocketServiceRoom)
	r.GET("rooms/ws", api.SocketServicePlatform)
	/* ---------------------------  Private routes  --------------------------- */

	// private := r.Group("/shoot")
	// public.Use(middleware.AuthMiddleware)

	// routes definition for finding and creating artists
	//public.GET("/friends", api.GetFriends)

	// start the server
	r.Run(":5000")
}
