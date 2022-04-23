package socket

import (
	"encoding/json"
	"log"
	"time"
)

const SendMessageAction = "send-message"
const JoinRoomAction = "join-room"
const LeaveRoomAction = "leave-room"
const UserJoinedAction = "user-join"
const UserLeftAction = "user-left"
const RoomJoinedAction = "room-joined"
const JoinRoomPrivateAction = "join-room-private"
const SocietyEvent = "society-event"
const GameEvent = "game-event"
type SocketEventMess struct {

}
type SocietyEventMess struct {
	SocketEventMess
	Content string `json:"content"`
}
type GameEventMess struct {
	SocketEventMess
	Action  string  `json:"action"`
}
type Message struct {
	Event  string  `json:"event"`
	Message SocketEventMess  `json:"message"`
	Target  *Room  `json:"target"`
	Sender  *Client `json:"sender"`
	CreatedAt time.Time `json:"created_at"`
}


func (message *Message) encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return json
}