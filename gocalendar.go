// Package gocalendar provides functionality for working with calendar in Go.
package gocalendar

import (
	"time"
)

// MonthCalendar returns matrix where each slice of integers is a week of month.
// Parameter weekStartsWith determines with which weekday the week starts.
// For example, if weekStartsWith is time.Sunday numbers in the slice will be arranged
// in the order corresponding to the order of the weekdays: Sunday, Monday, Tuesday, etc.
// In this way, slice [-1 -1 1 2 3 4 5] will mean that month will start in Tuesday.
// Similar to calendar.monthcalendar() in Python.
// Use year=0 to select the current year, and month=0 to select the current month.
func MonthCalendar(year int, month time.Month, weekStartsWith time.Weekday) [][]int {

	checkYearAndMonth(&year, &month)

	var daysInMonth = DaysInMonth(year, month)
	var weekdayOfFirstDayOfMonth = int(time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.Local).Weekday())
	var weekdayOfLastDayOfMonth = int(time.Date(year, time.Month(month), daysInMonth, 0, 0, 0, 0, time.Local).Weekday())

	switch {
	case int(weekStartsWith) == weekdayOfFirstDayOfMonth:
		weekdayOfFirstDayOfMonth = 0
	case int(weekStartsWith) > weekdayOfFirstDayOfMonth:
		weekdayOfFirstDayOfMonth = 7 - int(weekStartsWith) + weekdayOfFirstDayOfMonth
	case int(weekStartsWith) < weekdayOfFirstDayOfMonth:
		weekdayOfFirstDayOfMonth -= int(weekStartsWith)
	}

	switch {
	case int(weekStartsWith) == weekdayOfLastDayOfMonth:
		weekdayOfLastDayOfMonth = 0
	case int(weekStartsWith) > weekdayOfLastDayOfMonth:
		weekdayOfLastDayOfMonth = 7 - int(weekStartsWith) + weekdayOfLastDayOfMonth
	case int(weekStartsWith) < weekdayOfLastDayOfMonth:
		weekdayOfLastDayOfMonth -= int(weekStartsWith)
	}

	var weeksOfMonth [][]int
	var week []int
	for weekday := 0; weekday < weekdayOfFirstDayOfMonth; weekday++ {
		week = append(week, -1)
	}

	for day := 1; day <= daysInMonth; day++ {
		week = append(week, day)
		if (day+weekdayOfFirstDayOfMonth)%7 == 0 {
			weeksOfMonth = append(weeksOfMonth, week)
			week = []int{}
		}
	}

	for weekday := weekdayOfLastDayOfMonth; weekday < 6; weekday++ {
		week = append(week, -1)
	}
	if len(week) > 0 {
		weeksOfMonth = append(weeksOfMonth, week)
	}

	return weeksOfMonth

}

// DaysInMonth returns the number of days in the month.
// Use year=0 to select the current year,
// and month=0 to select the current month.
func DaysInMonth(year int, month time.Month) int {

	checkYearAndMonth(&year, &month)
	month++
	if month%12 > 0 {
		month %= 12
	}
	return time.Date(year, time.Month(month), 0, 0, 0, 0, 0, time.UTC).Day()

}

func checkYearAndMonth(year *int, month *time.Month) {

	if *year == 0 {
		*year = time.Now().Year()
	}
	if *month == 0 {
		*month = time.Now().Month()
	}

}
