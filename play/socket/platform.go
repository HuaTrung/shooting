package socket

//chatServer.go
type Platform struct {
	rooms      map[*Room]bool
	clients    map[*Client]bool
	Register   chan *Client
	unregister chan *Client
}

// NewWebsocketServer creates a new WsServer type
func NewPlatformServer() *Platform {
	return &Platform{
		clients:    make(map[*Client]bool),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		rooms:      make(map[*Room]bool),
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

func (server *Platform) unregisterClient(client *Client) {
	if _, ok := server.clients[client]; ok {
		delete(server.clients, client)
	}
}

func (server *Platform) registerClient(client *Client) {
	server.clients[client] = true
	room:=server.findRoomByID(client.IDRoom)
	if room!=nil {
		room.clients[client] = true
	}
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
	for client := range server.clients {
		client.send <- message
	}
}

func (server *Platform) findRoomByID(ID uint) *Room {
	var foundRoom *Room
	for room := range server.rooms {
		if room.GetId() == ID {
			foundRoom = room
			break
		}
	}

	return foundRoom
}

func (server *Platform) findUserByID(ID string)  {
	//var foundUser api.User
	for client,_  := range server.clients {
		if client.GetId() == ID {
			break
		}
	}

	//return foundUser
}
