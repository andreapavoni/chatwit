package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
  "log"
  "net/http"
	"text/template"
)

// flags
var addr = flag.String("addr", "localhost:8080", "http service address")

// templates
var chatTempl = template.Must(template.ParseFiles("chat.html"))
var indexTempl = template.Must(template.ParseFiles("index.html"))

// handlers
func chatHandler(c http.ResponseWriter, req *http.Request) {
	chatTempl.Execute(c, req.Host)
}

func indexHandler(c http.ResponseWriter, req *http.Request) {
	indexTempl.Execute(c, req.Host)
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

  http.HandleFunc("/", indexHandler)
  http.HandleFunc("/chat", chatHandler)
  http.Handle("/ws", websocket.Handler(wsHandler))

  if err := http.ListenAndServe(*addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}
