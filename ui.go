package main

import (
  "clengine"
  "strconv"
  "strings"
  "fmt"
  "os/exec"
  "os"
)

func Clear() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

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

func MakeHeader(w [][]clengine.Tile, name string) [][]clengine.Tile {
  var longest, current int
  var toReturn [][]clengine.Tile
  for i := range w {
    for j := range w[i] {
      current += len(w[i][j].Tile)
    }
    if current > longest {
      longest = current
    }
    current = 0
  }
  toReturn = append(toReturn, []clengine.Tile{clengine.Tile{Tile: strings.Repeat("-", longest)}})
  toReturn = append(toReturn, []clengine.Tile{clengine.Tile{Tile: name}})
  toReturn = append(toReturn, []clengine.Tile{clengine.Tile{Tile: strings.Repeat("-", longest)}})
  for i := range w {
    toReturn = append(toReturn, w[i])
  }
  return toReturn
}

func Feed(f [][]clengine.Tile, e Events) {
  var wtd [][]clengine.Tile
  var selection int
  var in string

  f = MakeHeader(f, "FEED")
  fmt.Println(len(f))

  for {
    wtd = clengine.DuplicateWorld(f)
    wtd[selection+3][1].BgColor = "white"
    wtd[selection+3][1].Color = "black"
    Clear()
    clengine.DrawCentered(wtd)

    fmt.Scanln(&in)
    switch strings.ToLower(in){
    case "w":
      if selection > 0 {
        selection--
      }
    case "s":
      if selection < len(f)-4 {
        selection++
      }
    case "q":
      Clear()
      os.Exit(0)
    case "e":
      ShowEvent(e.e[selection])
    }
  }
}

func ShowEvent(e Event) {
  var in string
  var w [][]clengine.Tile
  date := strconv.Itoa(e.Date.Day) + "." + strconv.Itoa(e.Date.Month) + "."
  w = append(w, make([]clengine.Tile, 0))
  w[0] = append(w[0], clengine.Tile{Tile: date})
  w = append(w, make([]clengine.Tile, 0))
  w[1] = append(w[1], clengine.Tile{Tile: e.Title})
  w = append(w, make([]clengine.Tile, 0))
  w[2] = append(w[2], clengine.Tile{Tile: e.Description})

  w = MakeHeader(w, "EVENT")
  Clear()
  clengine.DrawCentered(w)
  fmt.Scanln(&in)
  if strings.ToLower(in) == "q" {
    Clear()
    os.Exit(0)
  }
}
