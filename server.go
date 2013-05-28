package main

import (
  "code.google.com/p/go.net/websocket"
  "github.com/alloy-d/goauth"
  "github.com/gorilla/mux"
  "github.com/gorilla/sessions"
  "log"
  "net/http"
  "html/template"
  "github.com/shaoshing/train"
)

type Server struct {
  hub     *Hub
  oauth   *oauth.OAuth
  cookies *sessions.CookieStore
  tmpl map[string]*template.Template
}

type ConfigServer struct {
  oauthKey      string
  oauthSecret   string
  oauthCallback string

  storeSecret string
}

func NewServer(c *ConfigServer) *Server {
  s := Server{}

  s.tmpl = make(map[string]*template.Template)
  s.tmpl["index"] = NewTemplate("views/index.html", "views/layout.html")
  s.tmpl["chat"] = NewTemplate("views/chat.html", "views/layout.html")

  s.hub = NewHub(&s)

  s.oauth = NewTwitterOAuth(c.oauthKey, c.oauthSecret, c.oauthCallback)
  s.cookies = sessions.NewCookieStore([]byte(c.storeSecret))

  return &s
}

func (s *Server) Run(host string) {
  s.hub.run()

  router := mux.NewRouter()
  router.HandleFunc("/", s.homeHandler).Methods("GET")
  router.HandleFunc("/logout", s.logoutHandler).Methods("GET")
  router.HandleFunc("/auth/twitter", s.twitterAuthHandler).Methods("GET")
  router.HandleFunc("/auth/twitter/callback", s.twitterAuthCallbackHandler).Methods("GET")
  router.HandleFunc("/chat/{id:[A-Za-z0-9]+}", s.chatHandler).Methods("GET")
  router.Handle("/ws/{id:[A-Za-z0-9]+}", websocket.Handler(s.wsHandler))
  // this is no longer needed because we are using asset pipeline
  //router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))
  router.NotFoundHandler = http.HandlerFunc(s.notFound)


  http.Handle("/", router)

  // setup asset pipeline
  train.ConfigureHttpHandler(nil)
  train.Config.BundleAssets = true

  if err := http.ListenAndServe(host, nil); err != nil {
    log.Fatal("ListenAndServe:", err)
  }
}

func (s *Server) GetSession(req *http.Request, value string) string {
  if session, _ := s.cookies.Get(req, "session"); session.Values[value] != nil {
    return session.Values[value].(string)
  }
  return ""
}

// Handlers

type chatData struct {
  Host   string
  RoomId string
}

func (s *Server) homeHandler(c http.ResponseWriter, req *http.Request) {
  if nickname := s.GetSession(req, "user"); nickname == "" {
    s.tmpl["index"].ExecuteTemplate(c, "layout", req.Host)
  } else {
    // TODO: serve a dashboard-like page with connected friends and/or available chat rooms
    // based on friends and/or trending topics
    http.Redirect(c, req, ("/chat/" + nickname), 302)
  }
}

func (s *Server) chatHandler(c http.ResponseWriter, req *http.Request) {
  session, _ := s.cookies.Get(req, "session")

  if session.Values["user"] == nil {
    http.Redirect(c, req, "/", 403)
  }

  params := mux.Vars(req)
  roomId := params["id"]

  s.tmpl["chat"].ExecuteTemplate(c, "layout", &chatData{Host: req.Host, RoomId: roomId})
}

func (s *Server) notFound(c http.ResponseWriter, req *http.Request) {
  http.Redirect(c, req, "/", 302)
}

func (s *Server) wsHandler(ws *websocket.Conn) {
  params := mux.Vars(ws.Request())
  roomId := params["id"]
  nickname := s.GetSession(ws.Request(), "user")

  client := &Client{
    out:     make(chan *Command),
    ws:       ws,
    room:     roomId,
    hub:      s.hub,
    nickname: nickname,
  }

  s.hub.register <- client

  defer func() { s.hub.unregister <- client }()
  client.Run()
}

func (s *Server) logoutHandler(c http.ResponseWriter, req *http.Request) {
  if session, _ := s.cookies.Get(req, "session"); session.Values["user"] != nil {
    session.Options = &sessions.Options{MaxAge: -1}
    session.Save(req, c)
  }

  http.Redirect(c, req, "/", 302)
}
