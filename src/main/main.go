package main

import (
  "fmt"
  "github.com/marekmaskarinec/clengine"
  "os"
)

func main() {
  var e Events
  var in string
  var menu [][]clengine.Tile

  menu = MenuUI()

  e = Load()
  e.Sort()
  if len(os.Args) >= 2 {
  	switch os.Args[1]{
  	case "-c":
  		Clear()
  		clengine.DrawCentered(GetCal(), false)
  		fmt.Scanln()
  		Clear()
  	}
  } else {
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
	      case "c":
	      	Clear()
	        clengine.DrawCentered(GetCal(), false)
	        fmt.Scanln()
	      }
	  }
	}
}
