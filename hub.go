package main

import(
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

func newHub(server *Server) *Hub {
  return &Hub{
    broadcast:   make(chan *Message),
    register:    make(chan *Connection),
    unregister:  make(chan *Connection),
    rooms:       make(map[string]*Room),
    server: server,
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
            go c.ws.Close()
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
}

func (h *Hub) leaveRoom(c *Connection) {
  delete(h.rooms, c.room)
  close(c.send)
}

func (h *Hub) broadcastMessage(message *Message) {
  session, _ := h.server.store.Get(message.connection.ws.Request(), "session")
  nickname := session.Values["user"].(string)

  msg := fmt.Sprintf("%s -> %s", nickname, message.Text)
  message.Text = msg
  h.broadcast <- message
}
