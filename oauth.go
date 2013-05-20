package main


import (
  "fmt"
  "github.com/alloy-d/goauth"
  "net/http"
)

func newOAuth() *oauth.OAuth {
  o := new(oauth.OAuth)
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
