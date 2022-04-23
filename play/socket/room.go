package socket

import (
	"encoding/json"
	"log"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Room struct {
	Name string
	ID 	uint
	// Registered clients.
	clients map[*Client]bool

	// Number of people in room
	quantity int

	// Flag check is it full?
	full bool
	// Inbound messages from the clients.
	broadcast chan *Message

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	Private    bool `json:"private"`
}
// NewRoom creates a new Room
func NewRoom(server *Platform,name string, id uint) *Room{
	room:= &Room{
		Name:       name,
		ID: 		id,
		clients:    make(map[*Client]bool),
		full: 	false,
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
	go room.RunRoom()
	server.rooms[room] = true
	return room
}

// RunRoom runs our room, accepting various requests
func (room *Room) RunRoom() {
	bs, _ := json.Marshal(room.clients)
	log.Printf("Room with id=%d , name=%s %s",room.ID,room.Name,string(bs))
	for {
		select {

		case client := <-room.register:
			room.registerClientInRoom(client)

		case client := <-room.unregister:
			room.unregisterClientInRoom(client)

		case message := <-room.broadcast:
			room.broadcastToClientsInRoom(message.encode())
		}

	}
}

func (room *Room) registerClientInRoom(client *Client) {
	//room.notifyClientJoined(client)
	room.clients[client] = true
}

func (room *Room) unregisterClientInRoom(client *Client) {
	if _, ok := room.clients[client]; ok {
		delete(room.clients, client)
	}
}

func (room *Room) broadcastToClientsInRoom(message []byte) {
	for client := range room.clients {
		client.send <- message
	}
}
//
//func (room *Room) notifyClientJoined(client *Client) {
//	message := &Message{
//		Action:  SendMessageAction,
//		Target:  room,
//		Message: fmt.Sprintf("welcomeMessage", client.GetName()),
//	}
//
//	room.broadcastToClientsInRoom(message.encode())
//}

func (room *Room) GetId() uint {
	return room.ID
}

func (room *Room) GetName() string {
	return room.Name
}

func (room *Room) GetPrivate() bool {
	return room.Private
}