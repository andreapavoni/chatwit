package main

import (
  "code.google.com/p/go.net/websocket"
)

type connection struct {
  // The hub it belongs to
  hub *hub
  // The websocket connection.
  ws *websocket.Conn
  // Buffered channel of outbound messages.
  send chan string
  user string
}

func newConnection(hub *hub, ws *websocket.Conn, user string) *connection {
  return &connection {
    hub: hub,
    ws: ws,
    send: make(chan string, 256),
    user: user,
  }
}

func (c *connection) run() {
  go c.writer()
  c.reader()
}

func (c *connection) closeWs() {
  go c.ws.Close()
}

func (c *connection) reader() {
  defer func() { c.hub.unregisterConnection(c) }()
  for {
    var message string
    if err := websocket.Message.Receive(c.ws, &message); err != nil {
      break
    }
    c.hub.broadcastMessage(c.user, message)
  }
  c.closeWs()
}

func (c *connection) writer() {
  for message := range c.send {
    if err := websocket.Message.Send(c.ws, message); err != nil {
      break
    }
  }
  c.closeWs()
}

