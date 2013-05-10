package main

import "flag"

func main() {
  flag.Parse()
  addr := flag.String("addr", "localhost:8080", "http service address")
  server := newServer(*addr)
  server.run()
}

