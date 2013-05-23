package main

import (
  "code.google.com/p/go.net/websocket"
  "log"
)

type Client struct {
  // The websocket client.
  ws *websocket.Conn

  // Buffered channel of outbound messages.
  send chan string

  // the Room name
  room string

  // the Hub it belongs to
  hub *Hub

  nickname string
}

func (c *Client) reader() {
  for {
    var rcv string

    if err := websocket.Message.Receive(c.ws, &rcv); err != nil {
      log.Println("ERROR WS RCV: ", err)
      break
    }

    message := Command{Text: rcv, client: c}
    c.hub.broadcastMessage(&message)
  }
  c.ws.Close()
}

func (c *Client) writer() {
  for message := range c.send {
    if err := websocket.JSON.Send(c.ws, message); err != nil {
      break
    }
  }
  c.ws.Close()
}

// Listen for read/write messages
func (c *Client) Run() {
  go c.writer()
  c.reader()
}
