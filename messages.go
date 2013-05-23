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
