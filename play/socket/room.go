package socket

import (

	"fmt"
	"log"
	"shootingplane/database/repositories"
	"time"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Room struct {
	Name string
	ID 	uint
	// Registered clients.
	clients map[uint]*Client
	player_A *Client
	player_B *Client
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
func NewRoom(name string, id uint) *Room{
	room:= &Room{
		Name:       name,
		ID: 		id,
		clients:    make(map[uint]*Client),
		full: 		false,
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
	go room.RunRoom()
	return room
}

// RunRoom runs our room, accepting various requests
func (room *Room) RunRoom() {
	log.Printf("Room with id=%d , name=%s",room.ID,room.Name)
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
	if room.quantity<2 && room.full==false {
		room.clients[client.GetID()]=client
		room.quantity+=1
		room.notifyClientJoined(client)
		if room.quantity==2 {
			room.full=true
		}
	} else {
		room.notifyFullSlotsRoom(client)
	}
}

func (room *Room) notifyFullSlotsRoom(client *Client) {
	message := &Message{
		Event:  SocietyEvent,
		Target:  room.GetId(),
		Sender: 0,
		Message: &SocietyEventMess{
			Type: UnableToJoin,
			Content: fmt.Sprintf("This room is full."),
		},
		CreatedAt: time.Now(),
	}
	client.send <- message.encode()
}

func (room *Room) unregisterClientInRoom(client *Client) {
	if _, ok := room.clients[client.GetID()]; ok {
		delete(room.clients, client.GetID())
	}
	repositories.RemoveUserFromRoom(client.platform.db,room.ID,client.GetID())
	room.notifyClientLeave(client)
}

func (room *Room) broadcastToClientsInRoom(message []byte) {
	for id := range room.clients {
		room.clients[id].send <- message
	}
}
//
func (room *Room) notifyClientJoined(client *Client) {
	message := &Message{
		Event:  SocietyEvent,
		Target:  room.GetId(),
		Sender: 0,
		Message: &SocietyEventMess{
			Type: UserJoinedAction,
			Content: fmt.Sprintf("Welcome %s!",client.GetName()),
		},
		CreatedAt: time.Now(),
	}

	room.broadcastToClientsInRoom(message.encode())
}
func (room *Room) notifyClientLeave(client *Client) {
	message := &Message{
		Event:  SocietyEvent,
		Target:  room.GetId(),
		Sender: client.GetID(),
		Message: &SocietyEventMess{
			Type: UserLeftAction,
			Content: fmt.Sprintf("User %s! has left the room",client.GetName()),
		},
		CreatedAt: time.Now(),
	}

	room.broadcastToClientsInRoom(message.encode())
}
func (room *Room) GetId() uint {
	return room.ID
}

func (room *Room) GetName() string {
	return room.Name
}

func (room *Room) GetPrivate() bool {
	return room.Private
}