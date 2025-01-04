package ws

import (
	"log"
	"sync"

	"github.com/gofiber/contrib/websocket"
)

type WebSocketServer struct {
	clients   map[*websocket.Conn]bool
	path      string
	Broadcast chan Message
	mutex     sync.Mutex
}

type Message struct {
	Payload string `json:"payload"`
}

func NewWebSocketServer(path string) *WebSocketServer {
	return &WebSocketServer{
		clients:   make(map[*websocket.Conn]bool),
		Broadcast: make(chan Message),
		path:      path,
	}
}

func (server *WebSocketServer) AddClient(conn *websocket.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.clients[conn] = true
}

func (server *WebSocketServer) RemoveClient(conn *websocket.Conn) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	delete(server.clients, conn)
}

func (server *WebSocketServer) BroadcastMessage(message Message) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	for client := range server.clients {
		if err := client.WriteMessage(websocket.TextMessage, []byte(message.Payload)); err != nil {
			log.Println("Error broadcasting message:", err)
			client.Close()
			delete(server.clients, client)
		}
	}
}

func (server *WebSocketServer) HandleMessages() {
	for {
		message := <-server.Broadcast
		server.BroadcastMessage(message)
	}
}
