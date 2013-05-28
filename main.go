package main

import (
  "log"
  "flag"
)

func main() {
  file := flag.String("conf", "development.conf.json", "configuration file")
  flag.Parse()

  var config AppConfig

  if  err := ReadJSONConfig((*file), &config); err != nil {
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
