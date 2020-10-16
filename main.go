package main

import (
  "fmt"
  "clengine"
  "os"
)

func main() {
  var e Events
  var in string
  var menu [][]clengine.Tile

  menu = MenuUI()

  e = Load()
  e.Sort()

  for {
      Clear()
      clengine.DrawCentered(menu, true)
      fmt.Scanln(&in)
      switch in{
      case "f":
        Feed(GetFeed(e), e)
      case "n":
        AddEventUI(e)
      case "h":
        Help()
      case "q":
        Clear()
        os.Exit(0)
      }
  }
}
