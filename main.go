package main

import (
  "fmt"
)

func main() {
  var e Events
  var in string
  fmt.Scanln(&in)

  e.e = append(e.e, Event{Date: Day{Day: 16, Month: 10}, Title: "narozeniny"})
  e.e = append(e.e, Event{Date: Day{Day: 14, Month: 10}, Title: "neco jinyho"})
  fmt.Println(Save(e))

  e = Load()
  e.Sort()

  for {
      switch in{
      case "f":
        Feed(GetFeed(e))
      }
  }
}
