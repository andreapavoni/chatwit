package main

import (
  "code.google.com/p/go.net/websocket"
  "log"
)

type Connection struct {
  // The websocket connection.
  ws *websocket.Conn

  // Buffered channel of outbound messages.
  send chan string

  // the Room name
  room string

  // the Hub it belongs to
  hub *Hub

  nickname string
}

func (c *Connection) reader() {
  for {
    var rcv string

    if err := websocket.Message.Receive(c.ws, &rcv); err != nil {
      log.Println("ERROR WS RCV: ", err)
      break
    }

    message := Message{Text: rcv, connection: c}
    c.hub.broadcastMessage(&message)
  }
  c.ws.Close()
}

func (c *Connection) writer() {
  for message := range c.send {
    if err := websocket.JSON.Send(c.ws, message); err != nil {
      break
    }
  }
  c.ws.Close()
}

// Listen for read/write messages
func (c *Connection) Run() {
  go c.writer()
  c.reader()
}
