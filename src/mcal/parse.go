package main

import (
  "encoding/json"
  "io/ioutil"
)

func Save(toSave Events) error {
  tmp, err := json.Marshal(toSave.e)
  if err != nil {
    return err
  }
  ioutil.WriteFile("events.json", tmp, 0644)
  return nil
}

func Load() Events {
  var events Events
  dat, _ := ioutil.ReadFile("events.json")
  json.Unmarshal(dat, &events.e)

  return events
}
