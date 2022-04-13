package main

import (
	"errors"
	"flag"
	"fmt"
)

const (
	MONDAY    string = "monday"
	TUESDAY          = "tuesday"
	WEDNESDAY        = "wednesday"
	THURSDAY         = "thursday"
	FRIDAY           = "friday"
	SATURDAY         = "saturday"
	SUNDAY           = "sunday"
)

// Struct to encapsulate the names of the weekdays and
// possibly use them in functions
type WeekData struct {
	days [7]string
}

//this struct holds the data needed to start the print of the days
type MonthData struct {
	numberOfDays int
	startDay     string
	weekData     WeekData
}

//logic for the printing of the days
func (m MonthData) print() string {
	result := ""
	currentWeekday := m.startDay
	for day := 1; day <= m.numberOfDays; day++ {
		result = result + fmt.Sprintf("%02d", day) + " - " + currentWeekday[0:2] + "\n"
		var err error
		currentWeekday, err = m.weekData.next(currentWeekday)
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return result
}

//gets you the next day
func (w WeekData) next(current string) (string, error) {
	if current == SUNDAY {
		return MONDAY, nil
	}
	for index, day := range w.days {
		if day == current {
			return w.days[index+1], nil
		}
	}
	return "", errors.New("Day " + current + " not found")
}

func main() {
	//the weekdata struct is used in the month data struct
	wd := WeekData{[7]string{MONDAY, TUESDAY, WEDNESDAY, THURSDAY, FRIDAY, SATURDAY, SUNDAY}}

	//a number of month data structs mapped to the name of the months
	data := map[string]MonthData{
		"january":   MonthData{numberOfDays: 31, startDay: SATURDAY, weekData: wd},
		"febuary":   MonthData{numberOfDays: 28, startDay: TUESDAY, weekData: wd},
		"march":     MonthData{numberOfDays: 31, startDay: TUESDAY, weekData: wd},
		"april":     MonthData{numberOfDays: 30, startDay: FRIDAY, weekData: wd},
		"mai":       MonthData{numberOfDays: 31, startDay: SUNDAY, weekData: wd},
		"june":      MonthData{numberOfDays: 30, startDay: WEDNESDAY, weekData: wd},
		"july":      MonthData{numberOfDays: 31, startDay: FRIDAY, weekData: wd},
		"august":    MonthData{numberOfDays: 31, startDay: MONDAY, weekData: wd},
		"september": MonthData{numberOfDays: 30, startDay: THURSDAY, weekData: wd},
		"october":   MonthData{numberOfDays: 31, startDay: SATURDAY, weekData: wd},
		"november":  MonthData{numberOfDays: 30, startDay: TUESDAY, weekData: wd},
		"december":  MonthData{numberOfDays: 31, startDay: THURSDAY, weekData: wd},
	}

	month := flag.String("month", "default", "month to pick")
	flag.Parse()
	fmt.Println("----" + *month + "----")
	if *month != "" {
		fmt.Print(data[*month].print())
	}
}
