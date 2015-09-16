package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
)

var (
	timedate  string
	timezones string
)

func init() {
	flag.StringVar(&timedate, "time", "", "Time to convert, RFC1123 (Mon, 02 Jan 2006 15:04:05 MST)")
	flag.StringVar(&timezones, "zones", "UTC,CET,US/Pacific,US/Central,US/Eastern", "Timezone to convert time.")
}

func main() {
	flag.Parse()

	var t time.Time
	var err error
	if timedate == "" {
		t = time.Now()
	} else {
		t, err = time.Parse(time.RFC1123, timedate)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	zones := strings.Split(timezones, ",")

	fmt.Println(t.Format(time.RFC1123), "(Origin)")
	for _, zone := range zones {
		loc, err := time.LoadLocation(zone)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(t.In(loc).Format(time.RFC1123))
	}
}
