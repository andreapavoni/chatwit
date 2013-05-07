package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	"log"
	"net/http"
	"text/template"
)

var (
	//flags
	addr = flag.String("addr", "localhost:8080", "http service address")

	// templates
	chatTempl  = template.Must(template.ParseFiles("chat.html"))
	indexTempl = template.Must(template.ParseFiles("index.html"))

	// hub
	h = hub {
		broadcast:   make(chan string),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),
	}
)

// handlers
func chatHandler(c http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		var nick string

		if nick = req.FormValue("nickname"); nick != "" {
			http.SetCookie(c, &http.Cookie{Name: "user", Value: nick})
		} else {
			nick = randNickname()
		}

		http.SetCookie(c, &http.Cookie{Name: "user", Value: nick})

	} else { // it's a GET
		if _, err := req.Cookie("user"); err != nil {
			http.Redirect(c, req, "/", 403)
		}
	}

	chatTempl.Execute(c, req.Host)
}

func indexHandler(c http.ResponseWriter, req *http.Request) {
	indexTempl.Execute(c, req.Host)
}

func wsHandler(ws *websocket.Conn) {
	cookie, _ := ws.Request().Cookie("user")
	nick := cookie.Value

	c := &connection{send: make(chan string, 256), ws: ws, user: nick}
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
	http.Handle("/ws", websocket.Handler(wsHandler)) //GET

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
