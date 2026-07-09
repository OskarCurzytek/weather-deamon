package scheduler

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
)

func resubscribe(conn *websocket.Conn) {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		err := conn.WriteMessage(websocket.TextMessage, []byte(`{"a":111}`))
		if err != nil {
			fmt.Println("write error:", err)
			return
		}
		fmt.Println("Resubscribed")
	}
}
