package socket

import (
	"gorm.io/gorm"
	"log"
	"shootingplane/entity/models"
)

//chatServer.go
type Platform struct {
	rooms      map[*Room]bool
	clients    map[uint]*Client
	Register   chan *Client
	unregister chan *Client
	db *gorm.DB
}

// NewWebsocketServer creates a new WsServer type
func NewPlatformServer(Db *gorm.DB) *Platform {
	var rooms []models.Room
	result := Db.Table("personal.rooms").Unscoped().Find(&rooms)
	if result.Error != nil {
		log.Fatal("Cant locate any room.")
	}
	rooms_map:=make(map[*Room]bool)
	for i := 0; i < len(rooms); i++ {
		rooms_map[NewRoom(rooms[i].Name,rooms[i].ID)]=rooms[i].IsAcquired
	}
	return &Platform{
		clients:    make(map[uint]*Client),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      rooms_map,
		db: Db,
	}
}

// Run our websocket server, accepting various requests
func (server *Platform) Run() {

	for {
		select {

		case client := <-server.Register:
			server.registerClient(client)

		case client := <-server.unregister:
			server.unregisterClient(client)
		}

	}
}
func (server *Platform) ActiveRoom(room *Room) {
	log.Printf("Activate room %s and name %s",room.GetId(),room.GetName())
	r := server.FindRoomByID(room.ID)
	server.rooms[r] = true
}

func (server *Platform) DeactivateRoom(room *Room) {
	r := server.FindRoomByID(room.ID)
	server.rooms[r] = false
}

func (server *Platform) unregisterClient(client *Client) {
	if _, ok := server.clients[client.GetID()]; ok {
		delete(server.clients, client.GetID())
	}
}

func (server *Platform) registerClient(client *Client) {
	server.clients[client.GetID()] = client
}
//chatServer.go
func (server *Platform) findRoomByName(name string) *Room {
	var foundRoom *Room
	for room := range server.rooms {
		if room.GetName() == name {
			foundRoom = room
			break
		}
	}

	return foundRoom
}
func (server *Platform) broadcastToClients(message []byte) {
	for i := range server.clients {
		server.clients[i].send <- message
	}
}

func (server *Platform) FindRoomByID(ID uint) *Room {
	var foundRoom *Room
	for room := range server.rooms {
		if room.GetId() == ID {
			foundRoom = room
			break
		}
	}

	return foundRoom
}

func (server *Platform) findUserByID(ID uint) *Client {
	//var foundUser api.User
	if val, ok := server.clients[ID]; ok {
		return val
	} else {
		return nil
	}
	//return foundUser
}
