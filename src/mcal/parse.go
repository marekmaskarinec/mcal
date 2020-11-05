package main

import (
  "encoding/json"
  "io/ioutil"
  "os"
)

func Save(toSave Events) error {
  tmp, err := json.Marshal(toSave.e)

  err = ioutil.WriteFile(os.Getenv("HOME") + "/.config/mcal/events.json", tmp, 0644)

	if err != nil {
		return err
	}
	
  return nil
}

func Load() Events {
  var events Events
  dat, _ := ioutil.ReadFile(os.Getenv("HOME") + "/.config/mcal/events.json")
  json.Unmarshal(dat, &events.e)

  return events
}
