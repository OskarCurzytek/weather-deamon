package scheduler

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
	"mod.com/m/internal/decoder"
	"mod.com/m/internal/geo"
	"mod.com/m/internal/models"
)

var cityCoords = models.CityCoordinates{
	Lat:    50.0701406,
	Lon:    19.897822,
	Radius: 600,
}

func Loop() {
	u := url.URL{Scheme: "wss", Host: "ws7.blitzortung.org", Path: "/"}

	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	defer conn.Close()

	log.Println("connected")

	conn.WriteMessage(websocket.TextMessage, []byte(`{"a":111}`))

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
		if lightning != nil && geo.IsWithinRadius(cityCoords, lightning) {
			log.Println("lightning detected within radius:", lightning)
		} else {
			log.Println("lightning detected outside radius or failed to decode:", lightning)
		}
	}
}
