package main

// Events
const (
  MSG = iota
  JOIN
  PART
)

type Command struct {
  Event int
  Arguments interface{}
  client *Client
}

type MsgCommand struct {
  Body string
  Nickname string
}

type PartCommand struct {
  Nickname string
	Stats *RoomStats
}

type JoinCommand struct {
  Nickname string
	Stats *RoomStats
}
