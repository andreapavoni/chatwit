package main

type Room struct {
  Name string

	connections map[*Connection]bool
}
