package main

import (
  "log"
  "flag"
  "strings"
)

func main() {
  file := flag.String("conf", "development.conf.json", "configuration file")
  port := flag.String("port", "8080", "port to listen")
  flag.Parse()

  var conf AppConfig

  if  err := ReadJSONConfig((*file), &conf); err != nil {
    log.Fatal(err)
  }

  listen := conf.Address + ":" + *port

  srvConfig := ConfigServer{
    oauthKey:      conf.OAuth.Key,
    oauthSecret:   conf.OAuth.Secret,
    oauthCallback: strings.Replace(conf.OAuth.Callback, "ADDRESS", listen, -1),
    storeSecret:   conf.SessionSecret,
  }

  server := NewServer(&srvConfig)
  server.Run(listen)
}
