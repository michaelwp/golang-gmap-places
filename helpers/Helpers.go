package helpers

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func DateTimeFormat(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	locTime := t.In(loc)
	return locTime
}

func GenerateCode() (code string) {
	// create new seed
	rand.Seed(time.Now().UnixNano())
	// defined range number for random
	min := 100000
	max := 999999
	// random number
	randomNumber := strconv.Itoa(rand.Intn(max-min+1) + min)
	// return code
	return randomNumber
}

func DateFormat(d time.Time, l string) (date string, time string) {
	const (
		layoutUS  = "January 2, 2006'15:04:05"
		layoutISO = "2006-01-02'15:04:05"
	)

	var r string
	//var t time.Time

	if l == "US" {
		//t, _ := time.Parse(layoutUS, d)
		r = d.Format(layoutUS)
	} else if l == "ISO" {
		//t, _ := time.Parse(layoutISO, d)
		r = d.Format(layoutISO)
	}

	// split date string
	s := strings.Split(r, "'")

	//log.Println(t)
	return s[0], s[1]
}
