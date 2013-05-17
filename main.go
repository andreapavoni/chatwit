package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	"log"
	"net/http"
	"text/template"
  "github.com/gorilla/mux"
)

var indexTemplate = template.Must(template.ParseFiles("views/index.html"))
var chatTemplate = template.Must(template.ParseFiles("views/chat.html"))
var hub = newHub()

type chatData struct {
  Host string
  RoomId string
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	indexTemplate.Execute(c, req.Host)
}

func chatHandler(c http.ResponseWriter, req *http.Request) {
  if _, err := req.Cookie("user"); err != nil {
    http.Redirect(c, req, "/", 403)
  }

  params := mux.Vars(req)
  roomId := params["id"]

  chatTemplate.Execute(c, &chatData{Host: req.Host, RoomId: roomId})
}

func startChatHandler(c http.ResponseWriter, req *http.Request) {
  if nickname := req.FormValue("nickname"); nickname != "" {
    http.SetCookie(c, &http.Cookie{Name: "user", Value: nickname})
    http.Redirect(c, req, ("/chat/" + nickname), 302)
  } else {
    http.Redirect(c, req, "/", 403)
  }
}

func notFound(c http.ResponseWriter, req *http.Request) {
  http.Redirect(c, req, "/", 302)
}

func main() {
	flag.Parse()
  addr := flag.String("addr", "localhost:8080", "http service address")

  hub.run()

  router := mux.NewRouter()
  router.HandleFunc("/", homeHandler).Methods("GET")

  router.HandleFunc("/chat/{id:[A-Za-z0-9]+}", chatHandler).Methods("GET")
  router.HandleFunc("/chat", startChatHandler).Methods("POST")
  router.Handle("/ws/{id:[A-Za-z0-9]+}", websocket.Handler(wsHandler))
  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
  router.NotFoundHandler = http.HandlerFunc(notFound)

  http.Handle("/", router)

  if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
