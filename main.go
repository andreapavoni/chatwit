package main

import (
  "log"
)

func main() {
  var config AppConfig
  err := ReadJSONConfig("config.json", &config)

  if err != nil {
    log.Fatal(err)
  }

  srvConfig := ConfigServer{
    oauthKey:      config.OAuth.Key,
    oauthSecret:   config.OAuth.Secret,
    oauthCallback: config.OAuth.Callback,
    storeSecret:   config.SessionSecret,
  }

  server := NewServer(&srvConfig)
  server.Run(config.Address)
}
