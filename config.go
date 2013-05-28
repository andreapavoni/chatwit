package main

import (
  "encoding/json"
  "io/ioutil"
)

type AppConfig struct {
  Address string
  SessionSecret string
  OAuth struct {
    Key string
    Secret string
    Callback string
  }
}

func ReadJSONConfig(file string, data interface{}) (error) {
  config, err := ioutil.ReadFile(file)

  if err != nil {
    return err
  }

  if err := json.Unmarshal(config, &data); err != nil {
    return err
  }

  return nil
}
