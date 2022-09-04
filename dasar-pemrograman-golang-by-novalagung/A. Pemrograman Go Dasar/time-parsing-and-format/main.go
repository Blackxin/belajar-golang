package main

import (
	"fmt"
	"time"
)

func main() {
	// A.40.1. Penggunaan time.Time
	// time1 := time.Now()
	// fmt.Printf("time1: %v\n", time1)

	// time2 := time.Date(2002, 4, 1, 23, 0, 0, 0, time.Local)
	// fmt.Printf("time2: %v\n", time2)

	// A.40.2. Method Milik time.Time
	// now := time.Now()
	// fmt.Println("year:", now.Year(), "month:", now.Month(), "Weekday:", now.Weekday().String())
	// fmt.Println("Local:", now.Local(), "location:", now.Location())

	// A.40.3. Parsing dari string ke time.Time
	// var layoutFormat, value string
	// var date time.Time

	// layoutFormat = "2006-01-02 15:04:05"
	// value = "2015-09-02 08:04:00"
	// date, _ = time.Parse(layoutFormat, value)
	// fmt.Println(value, "\t->", date.String())

	// layoutFormat = "02/01/2006 MST"
	// value = "02/09/2015 WIB"
	// date, _ = time.Parse(layoutFormat, value)
	// fmt.Println(value, "\t\t->", date.String())

	// A.40.4. Predefined Layout Format Untuk Keperluan Parsing Time
	// date, _ := time.Parse(time.RFC822, "01 Apr 02 23:00 WIB")
	// fmt.Println(date.String())

	// A.40.5. Format dari time.Time ke string
	// date, _ := time.Parse(time.RFC822, "01 Apr 02 23:00 WIB")
	// dateS1 := date.Format("Monday, 02 January 2006 15:05 MST")
	// fmt.Println("dateS1", dateS1)

	// dateS2 := date.Format(time.RFC3339)
	// fmt.Println("dateS2", dateS2)

	// A.40.6. Handle Error Parsing time.Time
	date, err := time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(date)
}
