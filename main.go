package main

import (
  "flag"
  "fmt"
)

func main() {
  flag.Parse()

  // address
  address := flag.String("addr", "127.0.0.1:8080", "http service address")

  // oauth
  key := flag.String("oa-key", "M9MHfTfKDyF5yZM6xueTxg", "OAuth consumer key")
  secret := flag.String("oa-secret", "1lClcicoUNEKA1pycLLO0Jruo0NA2AgK3KhLFY4jo", "OAuth consumer secret")
  callbackUrl := fmt.Sprintf("http://%s/auth/twitter/callback", *address)

  // session cookie
  session := flag.String("session-secret", "foobarsecret", "Session secret key")

  config := ConfigServer{
    oauthKey:      *key,
    oauthSecret:   *secret,
    oauthCallback: callbackUrl,
    storeSecret:   *session,
  }

  server := NewServer(&config)
  server.Run(*address)
}
