package main

type Day struct {
  Day int
  Month int
}

func GetLm() map[int]int {
  lm := make(map[int]int)
  lm[1] = 31
  lm[2] = 28
  lm[3] = 31
  lm[4] = 30
  lm[5] = 31
  lm[6] = 30
  lm[7] = 31
  lm[8] = 31
  lm[9] = 30
  lm[10] = 31
  lm[11] = 30
  lm[12] = 31

  return lm
}

//checks if the date is valid
func (d *Day) CheckValidity() bool {
  if d.Month > 12 {
    return false
  }

  //map with lenghts of certain months
  lm := make(map[int]int)
  lm[1] = 31
  lm[2] = 28
  lm[3] = 31
  lm[4] = 30
  lm[5] = 31
  lm[6] = 30
  lm[7] = 31
  lm[8] = 31
  lm[9] = 30
  lm[10] = 31
  lm[11] = 30
  lm[12] = 31

  if lm[d.Month] < d.Day {
    return false
  }
  return true
}

type Event struct {
  Date Day
  Time string
  Title string
  Description string
}

type Events struct {
  e []Event
}
