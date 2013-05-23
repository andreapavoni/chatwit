package main

import (
  "fmt"
)

type Hub struct {
  // Registered rooms.
  rooms map[string]*Room

  // Inbound messages from the clients.
  broadcast chan *Command

  // Register requests from the clients.
  register chan *Client

  // Unregister requests from clients.
  unregister chan *Client

  server *Server
}

type Room struct {
  Name string

  clients map[*Client]bool
}

func NewHub(server *Server) *Hub {
  return &Hub{
    broadcast:  make(chan *Command),
    register:   make(chan *Client),
    unregister: make(chan *Client),
    rooms:      make(map[string]*Room),
    server:     server,
  }
}

func (h *Hub) run() {
  go func() {
    for {
      select {
      case c := <-h.register:
        h.joinRoom(c)

      case c := <-h.unregister:
        h.leaveRoom(c)

      case m := <-h.broadcast:
        for c := range h.rooms[m.client.room].clients {
          select {
          case c.send <- m.Text:
          default:
            h.leaveRoom(c)
          }
        }
      }
    }
  }()
}

func (h *Hub) joinRoom(c *Client) {
  var room *Room

  if room = h.rooms[c.room]; room == nil {
    room = &Room{Name: c.room, clients: make(map[*Client]bool)}
  }

  room.clients[c] = true
  h.rooms[c.room] = room

  h.statusMessage(("*** SERVER: " + c.nickname + " has joined " + c.room + " ****"), room)
}

func (h *Hub) leaveRoom(c *Client) {
  h.statusMessage(("*** SERVER: " + c.nickname + " has left " + c.room + " ****"), h.rooms[c.room])
  room := h.rooms[c.room]

  delete(room.clients, c)
  close(c.send)
  go c.ws.Close()

  // remove Room if empty
  if len(room.clients) == 0 {
    delete(h.rooms, c.room)
  }
}

// Broadcasts a user message to the Room
func (h *Hub) broadcastMessage(message *Command) {
  msg := fmt.Sprintf("%s -> %s", message.client.nickname, message.Text)
  message.Text = msg
  h.broadcast <- message
}

// Broadcasts Hub message to the Room
func (h *Hub) statusMessage(message string, room *Room) {
  for c := range room.clients {
    select {
    case c.send <- message:
    default:
      h.leaveRoom(c)
    }
  }
}
