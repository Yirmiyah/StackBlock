package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		// Read in a message
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		// Print out that message for clarity
		log.Printf("Received: %s", p)

		// Write message back to browser
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
