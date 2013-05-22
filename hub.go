package main

import (
  "fmt"
)

type Hub struct {
  // Registered rooms.
  rooms map[string]*Room

  // Inbound messages from the connections.
  broadcast chan *Message

  // Register requests from the connections.
  register chan *Connection

  // Unregister requests from connections.
  unregister chan *Connection

  server *Server
}

type Room struct {
  Name string

  connections map[*Connection]bool
}

type Message struct {
  // The text of the message
  Text string

  // Connection relative to the current Room
  connection *Connection
}

func newHub(server *Server) *Hub {
  return &Hub{
    broadcast:  make(chan *Message),
    register:   make(chan *Connection),
    unregister: make(chan *Connection),
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
        for c := range h.rooms[m.connection.room].connections {
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

func (h *Hub) joinRoom(c *Connection) {
  var room *Room

  if room = h.rooms[c.room]; room == nil {
    room = &Room{Name: c.room, connections: make(map[*Connection]bool)}
  }

  room.connections[c] = true
  h.rooms[c.room] = room

  h.statusMessage(("*** SERVER: " + c.nickname + " has joined " + c.room + " ****"), room)
}

func (h *Hub) leaveRoom(c *Connection) {
  h.statusMessage(("*** SERVER: " + c.nickname + " has left " + c.room + " ****"), h.rooms[c.room])
  room := h.rooms[c.room]

  delete(room.connections, c)
  close(c.send)
  go c.ws.Close()

  // remove Room if empty
  if len(room.connections) == 0 {
    delete(h.rooms, c.room)
  }
}

// Broadcasts a user message to the Room
func (h *Hub) broadcastMessage(message *Message) {
  msg := fmt.Sprintf("%s -> %s", message.connection.nickname, message.Text)
  message.Text = msg
  h.broadcast <- message
}

// Broadcasts Hub message to the Room
func (h *Hub) statusMessage(message string, room *Room) {
  for c := range room.connections {
    select {
    case c.send <- message:
    default:
      h.leaveRoom(c)
    }
  }
}
