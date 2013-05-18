package main

import (
  "code.google.com/p/go.net/websocket"
  "fmt"
  "github.com/gorilla/mux"
)

type Connection struct {
  // The websocket connection.
  ws *websocket.Conn

  // Buffered channel of outbound messages.
  send chan string

  // the Room name
  room string
}

type Message struct {
  // The text of the message
  Text string

  // Connection relative to the current Room
  connection *Connection
}

func (c *Connection) reader() {
  for {
    var text string
    message := Message{Text: text, connection: c}

    err := websocket.Message.Receive(c.ws, &message.Text)
    if err != nil {
      break
    }
    broadcastMessage(&message)
  }
  c.ws.Close()
}

func (c *Connection) writer() {
  for message := range c.send {
    err := websocket.Message.Send(c.ws, message)
    if err != nil {
      break
    }
  }
  c.ws.Close()
}

func wsHandler(ws *websocket.Conn) {
  params := mux.Vars(ws.Request())
  roomId := params["id"]

  c := &Connection{send: make(chan string, 256), ws: ws, room: roomId}

  hub.register <- c

  defer func() { hub.unregister <- c }()
  go c.writer()
  c.reader()
}

func broadcastMessage(message *Message) {
  cookie, _ := message.connection.ws.Request().Cookie("user")
  nickname := cookie.Value

  msg := fmt.Sprintf("%s -> %s", nickname, message.Text)
  message.Text = msg
  hub.broadcast <- message
}

