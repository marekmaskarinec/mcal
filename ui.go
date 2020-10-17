package main

import (
  "clengine"
  "strconv"
  "strings"
  "fmt"
  "os/exec"
  "os"
  "bufio"
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
    day += "."
    if e.e[i].Date.Day < 10 {
      day += " "
    }
    month = strconv.Itoa(e.e[i].Date.Month)
    month += "."
    if e.e[i].Date.Month < 10 {
      month += " "
    }
    w[i] = append(w[i], clengine.Tile{Tile: day + month + " | ", Color: "white", BgColor: "black"})
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
    clengine.DrawCentered(wtd, false)

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
    case "a":
      return
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
  clengine.DrawCentered(w, false)
  fmt.Scanln(&in)
  if strings.ToLower(in) == "q" {
    Clear()
    os.Exit(0)
  }
}

func AddEventUI(e Events) {
  reader := bufio.NewReader(os.Stdin)
  var w [][]clengine.Tile
  var in string
  var ne Event
  var err error
  w = append(w, make([]clengine.Tile, 0))
  w[0] = append(w[0], clengine.Tile{Tile: "description:"})
  w = MakeHeader(w, "ADD EVENT")

  w[3][0].Tile = "day:"
  Clear()
  clengine.DrawCentered(w, true)
  fmt.Scanln(&in)
  ne.Date.Day, err = strconv.Atoi(in)
  if err != nil {
    fmt.Println("bad date")
    return
  }

  w[3][0].Tile = "month:"
  Clear()
  clengine.DrawCentered(w, true)
  fmt.Scanln(&in)
  ne.Date.Month, err = strconv.Atoi(in)
  if err != nil {
    fmt.Println("bad date")
    return
  }

  if !ne.Date.CheckValidity() {
    fmt.Println("bad date")
    return
  }

  w[3][0].Tile = "title:"
  Clear()
  clengine.DrawCentered(w, true)
  in, _ = reader.ReadString('\n')
  in = strings.Replace(in, "\n", "", -1)
  ne.Title = in

  w[3][0].Tile = "description:"
  Clear()
  clengine.DrawCentered(w, true)
  in, _ = reader.ReadString('\n')
  in = strings.Replace(in, "\n", "", -1)
  ne.Description = in

  e.Add(ne)
}

func MenuUI() [][]clengine.Tile {
  var w [][]clengine.Tile

  w = append(w, make([]clengine.Tile, 0))
  w[0] = append(w[0], clengine.Tile{Tile: "Welcome to mcal. Friendly cli calendar"})
  w = append(w, make([]clengine.Tile, 0))
  w[1] = append(w[1], clengine.Tile{Tile: "Type `h` for help"})

  w = MakeHeader(w, "MCAL")
  return w
}

func Help() {
  var w [][]clengine.Tile

  w, _ = clengine.EditTile(w, clengine.V2(0,0), clengine.Tile{Tile: "n: add event"})
  w, _ = clengine.EditTile(w, clengine.V2(1,0), clengine.Tile{Tile: "f: feed"})
  w, _ = clengine.EditTile(w, clengine.V2(2,0), clengine.Tile{Tile: "q: quit"})

  w = MakeHeader(w, "MCAL")
  Clear()
  clengine.DrawCentered(w, false)
  fmt.Scanln()
}
