package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
  "log"
  "net/http"
	"text/template"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
var homeTempl = template.Must(template.ParseFiles("home.html"))

func homeHandler(c http.ResponseWriter, req *http.Request) {
	homeTempl.Execute(c, req.Host)
}

func wsHandler(ws *websocket.Conn) {
  c := &connection{send: make(chan string, 256), ws: ws, user: randInt()}
	h.register <- c
	defer func() { h.unregister <- c }()
	go c.writer()
	c.reader()
}

func main() {
	flag.Parse()

	go h.run()

  http.HandleFunc("/", homeHandler)
  http.Handle("/ws", websocket.Handler(wsHandler))

  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
