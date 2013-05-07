package main

import (
	"code.google.com/p/go.net/websocket"
  "fmt"
	"math/rand"
	"time"
)

type connection struct {
	// The websocket connection.
	ws *websocket.Conn

	// Buffered channel of outbound messages.
	send chan string

  user int
}

func (c *connection) reader() {
	for {
		var message string

		if err := websocket.Message.Receive(c.ws, &message); err != nil {
			break
		}

    msg := fmt.Sprintf("%d -> %s", c.user, message)

		h.broadcast <- msg
	}

	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {

		if err := websocket.Message.Send(c.ws, message); err != nil {
			break
		}
	}

	c.ws.Close()
}

func randInt() int {
	seed := time.Now().UTC().UnixNano()
	rand.Seed(seed)

	val := int(seed)
	if val < 0 {
		val = -val
	}

	return (1 + rand.Intn(val))
}
