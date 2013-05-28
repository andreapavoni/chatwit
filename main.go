package main

import (
  "log"
  "flag"
)

func main() {
  file := flag.String("conf", "development.conf.json", "configuration file")
  port := flag.String("port", "8080", "port to listen")
  flag.Parse()

  var conf AppConfig

  if  err := ReadJSONConfig((*file), &conf); err != nil {
    log.Fatal(err)
  }

  srvConfig := ConfigServer{
    oauthKey:      conf.OAuth.Key,
    oauthSecret:   conf.OAuth.Secret,
    oauthCallback: conf.OAuth.Callback,
    storeSecret:   conf.SessionSecret,
  }

  server := NewServer(&srvConfig)
  server.Run(conf.Address + ":" + *port)
}
