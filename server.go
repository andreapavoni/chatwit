package main

import (
  "code.google.com/p/go.net/websocket"
  "github.com/alloy-d/goauth"
  "github.com/gorilla/mux"
  "github.com/gorilla/sessions"
  "log"
  "net/http"
  "text/template"
)

type Server struct {
  hub     *Hub
  oauth   *oauth.OAuth
  cookies *sessions.CookieStore

  // TODO: these should go to map[string]*Template
  indexTemplate *template.Template
  chatTemplate  *template.Template
}

type ConfigServer struct {
  oauthKey      string
  oauthSecret   string
  oauthCallback string

  storeSecret string
}

func NewServer(c *ConfigServer) *Server {

  s := Server{}
  s.indexTemplate = template.Must(template.ParseFiles("views/index.html"))
  s.chatTemplate = template.Must(template.ParseFiles("views/chat.html"))
  s.hub = newHub(&s)

  // TODO: load these settings from command line flags
  s.oauth = NewTwitterOAuth(c.oauthKey, c.oauthSecret, c.oauthCallback)
  s.cookies = sessions.NewCookieStore([]byte(c.storeSecret))

  return &s
}

func (s *Server) Run(host string) {
  s.hub.run()

  router := mux.NewRouter()
  router.HandleFunc("/", s.homeHandler).Methods("GET")
  router.HandleFunc("/auth/twitter", s.twitterAuthHandler).Methods("GET")
  router.HandleFunc("/auth/twitter/callback", s.twitterAuthCallbackHandler).Methods("GET")
  router.HandleFunc("/chat/{id:[A-Za-z0-9]+}", s.chatHandler).Methods("GET")
  router.Handle("/ws/{id:[A-Za-z0-9]+}", websocket.Handler(s.wsHandler))
  router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
  router.NotFoundHandler = http.HandlerFunc(s.notFound)

  http.Handle("/", router)

  if err := http.ListenAndServe(host, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}

func (s *Server) GetSession(req *http.Request, value string) string {
  session, _ := s.cookies.Get(req, "session")
  return session.Values[value].(string)
}

// Handlers

type chatData struct {
  Host   string
  RoomId string
}

func (s *Server) homeHandler(c http.ResponseWriter, req *http.Request) {
  s.indexTemplate.Execute(c, req.Host)
}

func (s *Server) chatHandler(c http.ResponseWriter, req *http.Request) {
  session, _ := s.cookies.Get(req, "session")

  if session.Values["user"] == nil {
    http.Redirect(c, req, "/", 403)
  }

  params := mux.Vars(req)
  roomId := params["id"]

  s.chatTemplate.Execute(c, &chatData{Host: req.Host, RoomId: roomId})
}

func (s *Server) notFound(c http.ResponseWriter, req *http.Request) {
  http.Redirect(c, req, "/", 302)
}

func (s *Server) wsHandler(ws *websocket.Conn) {
  params := mux.Vars(ws.Request())
  roomId := params["id"]
  nickname := s.GetSession(ws.Request(), "user")

  conn := &Connection{
    send:     make(chan string, 256),
    ws:       ws,
    room:     roomId,
    hub:      s.hub,
    nickname: nickname,
  }

  s.hub.register <- conn

  defer func() { s.hub.unregister <- conn }()
  conn.Run()
}
