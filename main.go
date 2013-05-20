package main

import (
	"code.google.com/p/go.net/websocket"
	"flag"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"text/template"
)

var (
	indexTemplate = template.Must(template.ParseFiles("views/index.html"))
	chatTemplate  = template.Must(template.ParseFiles("views/chat.html"))
	store         = sessions.NewCookieStore([]byte("foobarquxsecret"))
	hub           = newHub()

	// TODO: load these settings from command line flags
	twitterOAuth = newTwitterOAuth("M9MHfTfKDyF5yZM6xueTxg", "1lClcicoUNEKA1pycLLO0Jruo0NA2AgK3KhLFY4jo", "http://127.0.0.1:8080/auth/twitter/callback")
)

type chatData struct {
	Host   string
	RoomId string
}

func homeHandler(c http.ResponseWriter, req *http.Request) {
	indexTemplate.Execute(c, req.Host)
}

func chatHandler(c http.ResponseWriter, req *http.Request) {
	session, _ := store.Get(req, "session")

	if session.Values["user"] == nil {
		http.Redirect(c, req, "/", 403)
	}

	params := mux.Vars(req)
	roomId := params["id"]

	chatTemplate.Execute(c, &chatData{Host: req.Host, RoomId: roomId})
}

func notFound(c http.ResponseWriter, req *http.Request) {
	http.Redirect(c, req, "/", 302)
}

func main() {
	flag.Parse()
	addr := flag.String("addr", "127.0.0.1:8080", "http service address")

	hub.run()

	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/auth/twitter", twitterAuthHandler).Methods("GET")
	router.HandleFunc("/auth/twitter/callback", twitterAuthCallbackHandler).Methods("GET")
	router.HandleFunc("/chat/{id:[A-Za-z0-9]+}", chatHandler).Methods("GET")
	router.Handle("/ws/{id:[A-Za-z0-9]+}", websocket.Handler(wsHandler))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
	router.NotFoundHandler = http.HandlerFunc(notFound)

	http.Handle("/", router)

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
