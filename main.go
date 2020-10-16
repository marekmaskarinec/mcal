package main

import (
  "fmt"
)

func main() {
  var e Events
  var in string

  e = Load()
  e.Sort()

  for {
      fmt.Scanln(&in)
      switch in{
      case "f":
        Feed(GetFeed(e), e)
      case "n":
        AddEventUI(e)
      }
  }
}
