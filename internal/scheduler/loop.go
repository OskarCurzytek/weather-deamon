package scheduler

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"mod.com/m/internal/decoder"
	"mod.com/m/internal/models"
)

func Loop(ch chan<- *models.Lightning) {
	u := url.URL{Scheme: "wss", Host: "ws7.blitzortung.org", Path: "/"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer conn.Close()

	log.Println("connected")

	err = conn.WriteMessage(websocket.TextMessage, []byte(`{"a":111}`))

	if err != nil {
		log.Println("subscribe error:", err)
		return
	}

	go resubscribe(conn)

	dec := decoder.NewDecoder()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}

		lightning := dec.Decode(msg)
		if lightning != nil {
			ch <- lightning
		}
	}
}
