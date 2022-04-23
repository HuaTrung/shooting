// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socket

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	ID       uuid.UUID `json:"id"`
	Name	string
	IDRoom  uint
	platform *Platform
	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte
}
func (client *Client) GetName() string {
	return client.Name
}

func (client *Client) GetID() uuid.UUID {
	return client.ID
}

func NewClient(conn *websocket.Conn, platform *Platform, name string, ID uint) *Client {
	client := &Client{
		Name:     name,
		conn:     conn,
		platform: platform,
		send:     make(chan []byte, 256),
		IDRoom:   ID,
	}

	return client
}
func (client *Client) ReadPump() {
	defer func() {
		client.disconnect()
	}()

	client.conn.SetReadLimit(maxMessageSize)
	client.conn.SetReadDeadline(time.Now().Add(pongWait))
	client.conn.SetPongHandler(func(string) error { client.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := client.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}
		client.handleNewMessage(jsonMessage)
	}

}

func (client *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		client.conn.Close()
	}()
	for {
		select {
		case message, ok := <-client.send:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The WsServer closed the channel.
				client.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := client.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Attach queued chat messages to the current websocket message.
			n := len(client.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-client.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (client *Client) disconnect() {
	client.platform.unregister <- client
	room:=client.platform.findRoomByID(client.IDRoom)
	room.unregister<-client
	close(client.send)
	client.conn.Close()
}
//
func (client *Client) handleNewMessage(jsonMessage []byte) {

	var message Message
	if err := json.Unmarshal(jsonMessage, &message); err != nil {
		log.Printf("Error on unmarshal JSON message %s", err)
		return
	}

	message.Sender = client

	switch message.Event {
	case GameEvent:
		roomID := message.Target.GetId()
		if room := client.platform.findRoomByID(roomID); room != nil {
			room.broadcast <- &message
		}

	case SocietyEvent:
		roomID := message.Target.GetId()
		if room := client.platform.findRoomByID(roomID); room != nil {
			room.broadcast <- &message
		}
	}

}
//
//func (client *Client) handleJoinRoomMessage(message Message) {
//	roomName := message.Message
//
//	client.joinRoom(roomName, nil)
//}
//
//func (client *Client) handleLeaveRoomMessage(message Message) {
//	room := client.platform.findRoomByID(message.Message)
//	if room == nil {
//		return
//	}
//
//	if _, ok := client.rooms[room]; ok {
//		delete(client.rooms, room)
//	}
//
//	room.unregister <- client
//}
//
//func (client *Client) handleJoinRoomPrivateMessage(message Message) {
//
//		target := client.platform.findUserByID(message.Message)
//
//	if target == nil {
//		return
//	}
//
//	// create unique room name combined to the two IDs
//	roomName := message.Message + client.ID.String()
//
//	// Join room
//	joinedRoom := client.joinRoom(roomName, target)
//
//	// Invite target user
//	if joinedRoom != nil {
//		client.inviteTargetUser(target, joinedRoom)
//	}
//
//}
//
//func (client *Client) joinRoom(roomName string, sender User) *Room {
//
//	room := client.platform.findRoomByName(roomName)
//	if room == nil {
//		room = client.platform.createRoom(roomName, sender != nil)
//	}
//
//	// Don't allow to join private rooms through public room message
//	if sender == nil && room.Private {
//		return nil
//	}
//
//	if !client.isInRoom(room) {
//
//		client.rooms[room] = true
//		room.register <- client
//
//		client.notifyRoomJoined(room, sender)
//	}
//
//	return room
//
//}
//
//func (client *Client) isInRoom(room *Room) bool {
//	if _, ok := client.rooms[room]; ok {
//		return true
//	}
//
//	return false
//}
//
//func (client *Client) inviteTargetUser(target models.User, room *Room) {
//	inviteMessage := &Message{
//		Action:  JoinRoomPrivateAction,
//		Message: target,
//		Target:  room,
//		Sender:  client,
//	}
//}
//
//func (client *Client) notifyRoomJoined(room *Room, sender models.User) {
//	message := Message{
//		Action: RoomJoinedAction,
//		Target: room,
//		Sender: sender,
//	}
//
//	client.send <- message.encode()
//}

func (client *Client) GetId() string {
	return client.ID.String()
}
