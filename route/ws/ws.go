package ws

import (
	"encoding/json"
	"log"

	"github.com/gofiber/contrib/websocket"
)

func ControlSkipHandler(conn *websocket.Conn, websocketServer *WebSocketServer) error {

	type Body struct {
		Skip bool `json:"skip"`
	}

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

		body := new(Body)
		err = json.Unmarshal(message, &body)
		if err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		bytes, _ := json.Marshal(body)

		websocketServer.Broadcast <- Message{
			Payload: string(bytes),
		}
	}

	return nil
}

func ControlSongEndHandler(conn *websocket.Conn, websocketServer *WebSocketServer) error {

	type Body struct {
		ID bool `json:"id"`
	}

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

		body := new(Body)
		err = json.Unmarshal(message, &body)
		if err != nil {
			log.Println("Unmarshal error:", err)
			continue
		}

		bytes, _ := json.Marshal(body)

		websocketServer.Broadcast <- Message{
			Payload: string(bytes),
		}
	}

	return nil
}
