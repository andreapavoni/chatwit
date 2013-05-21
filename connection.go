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
    c.hub.broadcastMessage(&message)
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

/*func (c *Connection) broadcastMessage(message *Message) {*/
  /*session, _ := c.server.store.Get(message.connection.ws.Request(), "session")*/
  /*nickname := session.Values["user"].(string)*/

  /*msg := fmt.Sprintf("%s -> %s", nickname, message.Text)*/
  /*message.Text = msg*/
  /*c.server.hub.broadcast <- message*/
/*}*/
