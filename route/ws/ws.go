package ws

import (
	"log"

	"github.com/gofiber/contrib/websocket"
)

func ControlSkipHandler(conn *websocket.Conn, websocketServer *WebSocketServer) error {
	defer func() {
		log.Println("Skip Client disconnected")
		websocketServer.RemoveClient(conn)
		conn.Close()
	}()

	log.Println("Skip Client connected")
	websocketServer.AddClient(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", message)

		websocketServer.Broadcast <- Message{
			Payload: string(message),
		}
	}

	return nil
}

func ControlSongEndHandler(conn *websocket.Conn, websocketServer *WebSocketServer) error {
	defer func() {
		log.Println("SongEnd Client disconnected")
		websocketServer.RemoveClient(conn)
		conn.Close()
	}()

	log.Println("SongEnd Client connected")
	websocketServer.AddClient(conn)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		log.Printf("Received: %s", message)

		websocketServer.Broadcast <- Message{
			Payload: string(message),
		}
	}

	return nil
}
