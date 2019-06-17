package gocalendar

import (
	"reflect"
	"testing"
	"time"
)

type MonthCalendarStruct struct {
	year           int
	month          time.Month
	weekStartsWith time.Weekday
	result         [][]int
}

var MonthCalendarTests = []MonthCalendarStruct{
	{
		year:           2019,
		month:          1, // January
		weekStartsWith: time.Sunday,
		result: [][]int{
			{-1, -1, 1, 2, 3, 4, 5},
			{6, 7, 8, 9, 10, 11, 12},
			{13, 14, 15, 16, 17, 18, 19},
			{20, 21, 22, 23, 24, 25, 26},
			{27, 28, 29, 30, 31, -1, -1},
		},
	},
	{
		year:           2019,
		month:          time.November,
		weekStartsWith: time.Monday,
		result: [][]int{
			{-1, -1, -1, -1, 1, 2, 3},
			{4, 5, 6, 7, 8, 9, 10},
			{11, 12, 13, 14, 15, 16, 17},
			{18, 19, 20, 21, 22, 23, 24},
			{25, 26, 27, 28, 29, 30, -1},
		},
	},
}

func TestMonthCalendar(t *testing.T) {

	for i, test := range MonthCalendarTests {

		res := MonthCalendar(test.year, test.month, test.weekStartsWith)
		if !reflect.DeepEqual(res, test.result) {
			t.Errorf(
				"\nTest No.%d\nYear: %d\nMonth: %v\nWeek starts with: %v\nExpected:%v\nGot:%v",
				i, test.year, test.month, test.weekStartsWith, test.result, res,
			)
		}

	}

}

type DaysInMonthStruct struct {
	year   int
	month  time.Month
	result int
}

var DaysInMonthTests = []DaysInMonthStruct{
	{
		year:   2019,
		month:  1, // January
		result: 31,
	},
	{
		year:   2019,
		month:  time.February,
		result: 28,
	},
	{
		year:   2000,
		month:  time.March,
		result: 31,
	},
	{
		year:   2016,
		month:  time.February,
		result: 29,
	},
	{
		year:   1997,
		month:  time.November,
		result: 30,
	},
}

func TestDaysInMonthCalendar(t *testing.T) {

	for i, test := range DaysInMonthTests {

		res := DaysInMonth(test.year, test.month)
		if !reflect.DeepEqual(res, test.result) {
			t.Errorf(
				"\nTest No.%d\nYear: %d\nMonth: %v\nExpected:%v\nGot:%v",
				i, test.year, test.month, test.result, res,
			)
		}

	}

}
