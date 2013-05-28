package main

import (
  "code.google.com/p/go.net/websocket"
)

type Client struct {
  // The websocket client.
  ws *websocket.Conn

  // Buffered channel of outbound messages.
  out chan *Command

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
      break
    }

    msg := MsgCommand{Body: rcv, Nickname: c.nickname}
    c.hub.broadcastMessage(&Command{Event: MSG, Arguments: msg, client: c})
  }
  c.ws.Close()
}

func (c *Client) writer() {
  for cmd := range c.out {
    if err := websocket.JSON.Send(c.ws, cmd); err != nil {
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
