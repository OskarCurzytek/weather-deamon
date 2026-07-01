package main

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "wss", Host: "ws7.blitzortung.org", Path: "/"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer conn.Close()

	log.Println("connected")

	conn.WriteMessage(websocket.TextMessage, []byte(`{"a":111}`))

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		log.Printf("type=%d, len=%d\n", msgType, len(msg))

		log.Println("message:", string(msg))
	}
}
