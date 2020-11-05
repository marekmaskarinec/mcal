package main

import (
  "github.com/marekmaskarinec/clengine"
  "strconv"
  "strings"
  "fmt"
  "os/exec"
  "os"
  "bufio"
  "bytes"
)

func Clear() {
  cmd := exec.Command("clear")
  cmd.Stdout = os.Stdout
  cmd.Run()
}
func Exec(comm, args, args2 string) string {
  cmd := exec.Command(comm, args, args2)
  cmd.Stdout = os.Stdout
  var out bytes.Buffer
  var stderr bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = &stderr
  err := cmd.Run()
  if err != nil {
    fmt.Println(stderr.String())
    panic(err)
  }

  return out.String()
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

  if len(e.e) == 0 {
  	f, _ = clengine.EditTile(f, clengine.V2(3, 0), clengine.Tile{Tile:"no events"})
  	fmt.Scanln()
  } else {
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
	      ShowEvent(e, selection)
	    case "a":
	      return
	    }
	  }
  }
}

func ShowEvent(e Events, index int) {
  var in string
  var w [][]clengine.Tile
  date := strconv.Itoa(e.e[index].Date.Day) + "." + strconv.Itoa(e.e[index].Date.Month) + "."
  w = append(w, make([]clengine.Tile, 0))
  w[0] = append(w[0], clengine.Tile{Tile: date})
  w = append(w, make([]clengine.Tile, 0))
  w[1] = append(w[1], clengine.Tile{Tile: e.e[index].Title})
  w = append(w, make([]clengine.Tile, 0))
  w[2] = append(w[2], clengine.Tile{Tile: e.e[index].Description})

  w = MakeHeader(w, "EVENT")
  Clear()
  clengine.DrawCentered(w, false)
  fmt.Scanln(&in)
  switch strings.ToLower(in) {
  case "q":
    Clear()
    os.Exit(0)
  case "r":
    e.Remove(index)
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

  w[3][0].Tile = "time:"
  Clear()
  clengine.DrawCentered(w, true)
  fmt.Println(&in)
  ne.Time = in


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

  w, _ = clengine.EditTile(w, clengine.V2(3,0), clengine.Tile{Tile: "c: month view"})
  w, _ = clengine.EditTile(w, clengine.V2(0,0), clengine.Tile{Tile: "n: add event"})
  w, _ = clengine.EditTile(w, clengine.V2(1,0), clengine.Tile{Tile: "f: feed"})
  w, _ = clengine.EditTile(w, clengine.V2(2,0), clengine.Tile{Tile: "q: quit"})

  w = MakeHeader(w, "MCAL")
  Clear()
  clengine.DrawCentered(w, false)
  fmt.Scanln()
}

func GetCal() [][]clengine.Tile {
  var w [][]clengine.Tile
  var lastWeekDay int
  weekDays := make(map[string]int)
  weekDays["Monday"] = 1
  weekDays["Tuesday"] = 2
  weekDays["Wednesday"] = 3
  weekDays["Thursday"] = 4
  weekDays["Friday"] = 5
  weekDays["Saturday"] = 6
  weekDays["Sunday"] = 7

  month, err := strconv.Atoi(strings.Split(Exec("date", "-u", "+%m"), "\n")[0])
  day, err := strconv.Atoi(strings.Split(Exec("date", "-u", "+%d"), "\n")[0])
  weekNumSt := ((day-1)/7)+2
  weekDaySt := weekDays[strings.Split(Exec("date", "-u", "+%A"), "\n")[0]]
  if err != nil {
    panic(err)
  }

  /*
  fmt.Println(month)
  fmt.Println(day)
  fmt.Println(weekNumSt)
  fmt.Println(weekDaySt)
  fmt.Scanln()
  */


  w, _ = clengine.EditTile(w, clengine.V2(0,0), clengine.Tile{Tile: "Mo "})
  w, _ = clengine.EditTile(w, clengine.V2(0,1), clengine.Tile{Tile: "Tu "})
  w, _ = clengine.EditTile(w, clengine.V2(0,2), clengine.Tile{Tile: "We "})
  w, _ = clengine.EditTile(w, clengine.V2(0,3), clengine.Tile{Tile: "Th "})
  w, _ = clengine.EditTile(w, clengine.V2(0,4), clengine.Tile{Tile: "Fr "})
  w, _ = clengine.EditTile(w, clengine.V2(0,5), clengine.Tile{Tile: "Sa "})
  w, _ = clengine.EditTile(w, clengine.V2(0,6), clengine.Tile{Tile: "Su"})

  weekNum := weekNumSt
  weekDay := weekDaySt+1
  for i:=day+1; i != 0; i-- {
    if i >= 10 {
      w, _ = clengine.EditTile(w, clengine.V2(weekNum, weekDay), clengine.Tile{Tile: strconv.Itoa(i) + " "})
    } else {
      w, _ = clengine.EditTile(w, clengine.V2(weekNum, weekDay), clengine.Tile{Tile: strconv.Itoa(i) + "  "})
    }
    lastWeekDay = weekDay
    if weekDay > 1 {
      weekDay--
    } else if weekNum >= 1 {
      weekNum--
      weekDay = 7
    } else {
      break
    }
  }

  for i:=lastWeekDay-1; i!=0; i-- {
    w, _ = clengine.EditTile(w, clengine.V2(1, i), clengine.Tile{Tile: "   "})
  }

  weekNum = weekNumSt
  weekDay = weekDaySt
  for i:=day; i<=GetLm()[month]; i++ {
	if i >= 10 {
      w, _ = clengine.EditTile(w, clengine.V2(weekNum, weekDay), clengine.Tile{Tile: strconv.Itoa(i) + " "})
    } else {
      w, _ = clengine.EditTile(w, clengine.V2(weekNum, weekDay), clengine.Tile{Tile: strconv.Itoa(i) + "  "})
    }  	
	if weekDay == 7 {
      weekDay = 1
      weekNum++
    } else {
      weekDay++
    }
  }
  w[weekNumSt][weekDaySt].BgColor = "cyan"
  w = MakeHeader(w, "MONTH")

  return w
}

func (e *Events) Cal(w [][]clengine.Tile) {
  var in string
  wtd := clengine.DuplicateWorld(w)
  cords := clengine.V2(0, 0)

  for {
    wtd = clengine.DuplicateWorld(w)
    wtd[cords.X+4][cords.Y].BgColor = "white"
    wtd[cords.X+4][cords.Y].Color = "black"
    Clear()
    clengine.DrawCentered(wtd, false)
	fmt.Println(cords)
    fmt.Scanln(&in)

    switch strings.ToLower(in) {
    case "s":
      if cords.X < len(w)-1 {
        cords.X++
      }
    case "w":
      if cords.X > 0 {
        cords.X--
      }
    case "a":
      if cords.Y > 0 {
        cords.Y--
      }
    case "d":
      if cords.Y < len(w[0])-1 {
        cords.Y++
      }
    }

  }
}
