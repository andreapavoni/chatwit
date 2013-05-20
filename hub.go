package main

type Hub struct {
  // Registered rooms.
  rooms map[string]*Room

  // Inbound messages from the connections.
  broadcast chan *Message

  // Register requests from the connections.
  register chan *Connection

  // Unregister requests from connections.
  unregister chan *Connection
}

type Room struct {
  Name string

  connections map[*Connection]bool
}

func newHub() *Hub {
  return &Hub{
    broadcast:   make(chan *Message),
    register:    make(chan *Connection),
    unregister:  make(chan *Connection),
    rooms:       make(map[string]*Room),
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

  if room = hub.rooms[c.room]; room == nil {
    room = &Room{Name: c.room, connections: make(map[*Connection]bool)}
  }

  room.connections[c] = true
  hub.rooms[c.room] = room
}

func (h *Hub) leaveRoom(c *Connection) {
  delete(hub.rooms, c.room)
  close(c.send)
}

