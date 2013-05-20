package main


import (
  "log"
  "github.com/alloy-d/goauth"
  "net/http"
)

func newTwitterOAuth(key, secret, callback string) *oauth.OAuth {
  o := new(oauth.OAuth)
  o.ConsumerKey = key
  o.ConsumerSecret = secret
  o.Callback = callback
  o.RequestTokenURL = "https://api.twitter.com/oauth/request_token"
  o.OwnerAuthURL = "https://api.twitter.com/oauth/authorize"
  o.AccessTokenURL = "https://api.twitter.com/oauth/access_token"
  o.SignatureMethod = "HMAC-SHA1"
  return o
}

func twitterAuthHandler(c http.ResponseWriter, req *http.Request) {
  err := twitterOAuth.GetRequestToken()
  if err != nil {
    log.Println(err)
    return
  }

  url, err := twitterOAuth.AuthorizationURL()
  if err != nil {
    log.Println(err)
    return
  }

  session, _ := store.Get(req, "session")
  session.Values["requestToken"] = twitterOAuth.RequestToken
  session.Values["requestSecret"] = twitterOAuth.RequestSecret
  session.Save(req, c)

  http.Redirect(c, req, url, 302)
}

func twitterAuthCallbackHandler(c http.ResponseWriter, req *http.Request) {
  session, _ := store.Get(req, "session")
  twitterOAuth.RequestToken = session.Values["requestToken"].(string)
  twitterOAuth.RequestSecret = session.Values["requestSecret"].(string)

  req.ParseForm()
  token := req.Form.Get("oauth_verifier")

  err := twitterOAuth.GetAccessToken(token)

  if err != nil {
    log.Println(err)
    http.Redirect(c, req, "/", 403)
    return
  }

  session.Values["user"] = twitterOAuth.UserName()
  session.Save(req, c)

  http.Redirect(c, req, ("/chat/" + twitterOAuth.UserName()), 302)
}
