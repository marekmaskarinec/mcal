package main

func (e *Events) Sort() {
  var tmp Event

  for i:=0; i<len(e.e)-1; i++ {
    for j:=0; j<len(e.e)-1; j++ {
      if e.e[j].Date.Month > e.e[j+1].Date.Month {
        tmp = e.e[j]
        e.e[j] = e.e[j+1]
        e.e[j+1] = tmp
      } else if e.e[j].Date.Month == e.e[j+1].Date.Month && e.e[j].Date.Day > e.e[j+1].Date.Day {
        tmp = e.e[j]
        e.e[j] = e.e[j+1]
        e.e[j+1] = tmp
      }
    }
  }
}
