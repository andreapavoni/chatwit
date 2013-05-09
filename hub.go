package main

import (
  "fmt"
  "code.google.com/p/go.net/websocket"
)

type hub struct {
  // Registered connections.
  connections map[*connection]bool

  // Inbound messages from the connections.
  broadcast chan string

  // Register requests from the connections.
  register chan *connection

  // Unregister requests from connections.
  unregister chan *connection
}

func newHub() *hub {
  return &hub {
    broadcast: make(chan string),
    register: make(chan *connection),
    unregister: make(chan *connection),
    connections: make(map[*connection]bool),
  }
}

func (h *hub) run() {
  for {
    select {
    case c := <-h.register:
      h.connections[c] = true
    case c := <-h.unregister:
      h.removeConnection(c)
    case m := <-h.broadcast:
      for c := range h.connections {
        select {
        case c.send <- m:
        default:
          h.removeConnection(c)
          c.closeWs()
        }
      }
    }
  }
}

func (h *hub) removeConnection(c *connection) {
  delete(h.connections, c)
  close(c.send)
}

func (h *hub) registerConnection(nick string, ws *websocket.Conn) *connection {
  c := newConnection(h, ws, nick)
  h.register <- c
  return c
}

func (h *hub) unregisterConnection(c *connection) {
  h.unregister <- c
}

func (h *hub) broadcastMessage(user string, message string) {
  msg := fmt.Sprintf("%s -> %s", user, message)
  h.broadcast <- msg
}

