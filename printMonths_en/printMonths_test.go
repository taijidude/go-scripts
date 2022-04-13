package main

import (
	"testing"
)

func TestNext(t *testing.T) {
	wd := WeekData{[7]string{MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY}}

	nextDay, _ := wd.next(MONDAY)

	if nextDay != TUESDAY {
		t.Errorf("Expected '"+TUESDAY+"', got '%s'", nextDay)
	}
}

func TestPrint(t *testing.T) {

	wd := WeekData{[7]string{MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY}}
	md := MonthData{numberOfDays: 1, startDay: SATURDAY, weekData: wd}

	result := md.print()
	if result != "01 - sa\n" {
		t.Errorf("Expected '01 - sa\n', got '%s'", result)
	}

	md = MonthData{numberOfDays: 3, startDay: SATURDAY, weekData: wd}
	result = md.print()
	if result != "01 - sa\n02 - su\n03 - mo\n" {
		t.Errorf("Expected '01 - sa\n02 - su\n03 - mo\n', got '%s'", result)
	}
}
