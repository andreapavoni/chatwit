package main

import (
  "code.google.com/p/go.net/websocket"
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
    var text string
    message := Message{Text: text, connection: c}

    if err := websocket.Message.Receive(c.ws, &message.Text); err != nil {
      break
    }
    c.hub.broadcastMessage(&message)
  }
  c.ws.Close()
}

func (c *Connection) writer() {
  for message := range c.send {
    if err := websocket.Message.Send(c.ws, message); err != nil {
      break
    }
  }
  c.ws.Close()
}
