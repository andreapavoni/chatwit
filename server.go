package main

import (
  "code.google.com/p/go.net/websocket"
  "fmt"
  "log"
  "math/rand"
  "net/http"
  "text/template"
  "time"
  "github.com/gorilla/mux"
)

func randNickname() string {
  seed := time.Now().UTC().UnixNano()
  rand.Seed(seed)

  val := int(seed)
  if val < 0 {
    val = -val
  }

  return fmt.Sprintf("anonymous-%d", (1 + rand.Intn(val)))
}

// handlers

type server struct {
  chatTemplate *template.Template
  indexTemplate *template.Template
  hub *hub
  addr string
}

func newServer(addr string) *server {
  return &server {
    chatTemplate: template.Must(template.ParseFiles("chat.html")),
    indexTemplate: template.Must(template.ParseFiles("index.html")),
    hub: newHub(),
    addr: addr,
  }
}

func (s *server) run() {
  s.hub.run()

  router := mux.NewRouter()

  /* router.HandleFunc("/", ProductsHandler).Methods("GET")*/

  router.HandleFunc("/", func (c http.ResponseWriter, req *http.Request) {
    s.indexTemplate.Execute(c, req.Host)
  }).Methods("GET")

  router.HandleFunc("/chat", func (c http.ResponseWriter, req *http.Request) {
    if _, err := req.Cookie("user"); err != nil {
      http.Redirect(c, req, "/", 403)
    }

    s.chatTemplate.Execute(c, req.Host)

  }).Methods("GET")


  router.HandleFunc("/chat", func (c http.ResponseWriter, req *http.Request) {
    var nick string

    if nick = req.FormValue("nickname"); nick != "" {
      http.SetCookie(c, &http.Cookie{Name: "user", Value: nick})
    } else {
      nick = randNickname()
    }

    http.SetCookie(c, &http.Cookie{Name: "user", Value: nick})
    s.chatTemplate.Execute(c, req.Host)

  }).Methods("POST")

  router.Handle("/ws", websocket.Handler(func (ws *websocket.Conn) {
    cookie, _ := ws.Request().Cookie("user")
    nick := cookie.Value
    s.hub.registerConnection(nick, ws)
  }))

  http.Handle("/", router)

  if err := http.ListenAndServe(s.addr, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}

