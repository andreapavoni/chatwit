package main

import (
  "code.google.com/p/go.net/websocket"
  "flag"
  "log"
  "net/http"
  "text/template"
  "github.com/gorilla/mux"
  "github.com/gorilla/sessions"
  "github.com/alloy-d/goauth"
  "fmt"
)

var indexTemplate = template.Must(template.ParseFiles("views/index.html"))
var chatTemplate = template.Must(template.ParseFiles("views/chat.html"))
var store = sessions.NewCookieStore([]byte("foobarquxsecret"))
var hub = newHub()

type chatData struct {
  Host string
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

func newOAuth() *oauth.OAuth {
  o := new (oauth.OAuth)
  o.ConsumerKey = "M9MHfTfKDyF5yZM6xueTxg"
  o.ConsumerSecret = "1lClcicoUNEKA1pycLLO0Jruo0NA2AgK3KhLFY4jo"
  o.Callback = "http://127.0.0.1:8080/auth/twitter/callback"
  o.RequestTokenURL = "https://api.twitter.com/oauth/request_token"
  o.OwnerAuthURL = "https://api.twitter.com/oauth/authorize"
  o.AccessTokenURL = "https://api.twitter.com/oauth/access_token"
  o.SignatureMethod = "HMAC-SHA1"
  return o
}

func twitterAuthHandler(c http.ResponseWriter, req *http.Request) {
  o := newOAuth()

  err := o.GetRequestToken()
  if err != nil {
    fmt.Println(err)
    return 
  }

  url, err := o.AuthorizationURL()
  if err != nil { 
    fmt.Println(err)
    return 
  }

  session, _ := store.Get(req, "session")
  session.Values["requestToken"] = o.RequestToken
  session.Values["requestSecret"] = o.RequestSecret
  session.Save(req, c)

  http.Redirect(c, req, url, 302)
}

func twitterAuthCallbackHandler(c http.ResponseWriter, req *http.Request) {
  o := newOAuth()

  session, _ := store.Get(req, "session")
  o.RequestToken = session.Values["requestToken"].(string)
  o.RequestSecret = session.Values["requestSecret"].(string)

  req.ParseForm()
  token := req.Form.Get("oauth_verifier")

  err := o.GetAccessToken(token)
  if err != nil {
    fmt.Println(err)
    http.Redirect(c, req, "/", 403)
    return
  }

  session.Values["user"] = o.UserName()
  session.Save(req, c)

  http.Redirect(c, req, ("/chat/" + o.UserName()), 302)
}

func main() {
  flag.Parse()
  addr := flag.String("addr", "localhost:8080", "http service address")

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

