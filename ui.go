package main

import (
  "clengine"
  "strconv"
  "fmt"
  "os/exec"
)

func GetFeed(e Events) [][]clengine.Tile {
  var w [][]clengine.Tile
  var day, month string

  for i := range e.e {
    w = append(w, make([]clengine.Tile, 0))
    day = strconv.Itoa(e.e[i].Date.Day)
    month = strconv.Itoa(e.e[i].Date.Month)
    w[i] = append(w[i], clengine.Tile{Tile: day + "." + month + ". | ", Color: "white", BgColor: "black"})
    w[i] = append(w[i], clengine.Tile{Tile: e.e[i].Title, Color: "white", BgColor: "black"})
  }

  return w
}

func Feed(f [][]clengine.Tile, e Events) {
  var wtd [][]clengine.Tile
  var selection int
  var in string
  clengine.DrawWorld(f)
  for {
    fmt.Scanln(&in)
    switch in{
    case "w":
      if selection > 0 {
        selection--
      }
    case "s":
      if selection < len(f)-1 {
        selection++
      }
    case "q":
      break
      return
    case "e":
      ShowEvent(e.e[selection])
    }

    wtd = clengine.DuplicateWorld(f)
    wtd[selection][1].BgColor = "white"
    wtd[selection][1].Color = "black"
    exec.Command("clear").Run()
    fmt.Println("FEED")
    fmt.Println("____")
    clengine.DrawWorld(wtd)
  }
}

func ShowEvent(e Event) {
  fmt.Println("EVENT")
  fmt.Println("_____")
  fmt.Printf("%d.%d.\n", e.Date.Day, e.Date.Month)
  fmt.Println(e.Title)
  fmt.Println(e.Description)
}
