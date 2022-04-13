package main

import (
	"testing"
)

func TestNext(t *testing.T) {
	wd := WeekData{[7]string{"montag", "dienstag", "mittwoch", "donnerstag", "freitag", "samstag", "sonntag"}}

	nextDay, _ := wd.next("montag")

	if nextDay != "dienstag" {
		t.Errorf("Expected 'dienstag', got '%s'", nextDay)
	}
}

func TestPrint(t *testing.T) {

	wd := WeekData{[7]string{"montag", "dienstag", "mittwoch", "donnerstag", "freitag", "samstag", "sonntag"}}
	md := MonthData{numberOfDays: 1, startDay: "samstag", weekData: wd}

	result := md.print()
	if result != "01 - sa\n" {
		t.Errorf("Expected '01 samstag\n', got '%s'", result)
	}

	md = MonthData{numberOfDays: 3, startDay: "samstag", weekData: wd}
	result = md.print()
	if result != "01 - sa\n02 - so\n03 - mo\n" {
		t.Errorf("Expected '01 - sa\n02 - so\n03 - mo\n', got '%s'", result)
	}
}
