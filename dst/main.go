package main

import (
	"fmt"
	"time"
)

var datePatterns = []string{
	"2006-01-02T15:04:05Z", //should always be this one, but seeing others in sandbox tests?
	"2006-01-02T15:04:05-07:00",
	"2006-01-02T15:04:05.000-07:00",
}

var serviceTimeZone *time.Location

func parseDateString(dateString string) (time.Time, string, error) {
	var datePattern string
	var startDate time.Time
	var err error
	for _, pattern := range datePatterns {
		datePattern = pattern
		startDate, err = time.ParseInLocation(datePattern, dateString, serviceTimeZone)
		if err == nil {
			break
		}
	}
	return startDate, datePattern, err
}

func main() {
	serviceTimeZone, _ = time.LoadLocation("US/Pacific")
	dates := []string{
		"2017-10-25T01:30:05-07:00",
		"2017-10-25T01:30:05Z",
		"2017-11-01T01:30:05-07:00",
		"2017-11-01T01:30:05Z",
		"2017-11-06T01:30:05-07:00",
		"2017-11-06T01:30:05Z",
	}
	for _, startDateString := range dates {
		startDate, datePattern, err := parseDateString(startDateString)
		if err != nil {
			fmt.Println(err)
			continue
		}
		dueDate := time.Date(
			startDate.Year(),
			startDate.Month(),
			startDate.Day(),
			12, 0, 0, 0, // set time to 12 Noon
			startDate.Location())
		for dueDate.Weekday() != time.Sunday {
			dueDate = dueDate.AddDate(0, 0, 1)
		}
		dueDateString := dueDate.Format(datePattern)
		fmt.Println(startDateString, "->", dueDateString)
	}
}
