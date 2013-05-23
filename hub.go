package main

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
      case client := <-h.register:
        h.joinRoom(client)

      case client := <-h.unregister:
        h.leaveRoom(client)

      case cmd := <-h.broadcast:
        for c := range h.rooms[cmd.client.room].clients {
          select {
          case c.out <- cmd:
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

  join := JoinCommand{Nickname: c.nickname}
  h.statusMessage(&Command{Event: JOIN, Arguments: join, client: c})
}

func (h *Hub) leaveRoom(c *Client) {
  part := PartCommand{Nickname: c.nickname}
  room := h.rooms[c.room]

  h.statusMessage(&Command{Event: PART, Arguments: part, client: c})

  delete(room.clients, c)
  close(c.out)
  go c.ws.Close()

  // remove Room if empty
  if len(room.clients) == 0 {
    delete(h.rooms, c.room)
  }
}

// Broadcasts a user message to the Room
func (h *Hub) broadcastMessage(cmd *Command) {
  h.broadcast <- cmd
}

// Broadcasts Hub message to the Room
func (h *Hub) statusMessage(cmd *Command) {
  for c := range cmd.client.hub.rooms[cmd.client.room].clients {
    select {
    case c.out <- cmd:
    default:
    }
  }
}
