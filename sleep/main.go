package main

import (
	"flag"
	"log"
	"time"
)

var expectedFormat = "2006-01-02"

// parseTime validates and parses a given date string.
func parseTime(target string) (time.Time, error) {
  r, err := time.Parse(time.DateOnly, target)
  if err != nil {
    return time.Now(), err 
  }

  return r, nil
}

// calcSleeps returns the number of sleeps until the target.
func calcSleeps(target time.Time) float64 {
  return time.Until(target).Hours() / 24
}

func main() {
	bday := flag.String("bday", "", "Your next bday in YYYY-MM-DD format")
	flag.Parse()

	target, err := parseTime(*bday)
  if err != nil {
    log.Fatalf("Invalid date, your input must in the format of YYYY-MM-DD")
  }

	log.Printf("You have %d sleeps until your birthday. Hurray!",
		int(calcSleeps(target)))
}
