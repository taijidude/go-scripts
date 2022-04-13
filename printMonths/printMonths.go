package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type WeekData struct {
	days [7]string
}

type MonthData struct {
	numberOfDays int
	startDay     string
	weekData     WeekData
}

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

func (w WeekData) next(current string) (string, error) {
	if current == "sonntag" {
		return "montag", nil
	}
	for index, day := range w.days {
		if day == current {
			return w.days[index+1], nil
		}
	}
	return "", errors.New("Kein Tag gefunden")
}

func main() {
	wd := WeekData{[7]string{"montag", "dienstag", "mittwoch", "donnerstag", "freitag", "samstag", "sonntag"}}

	data := map[string]MonthData{
		"januar":    MonthData{numberOfDays: 31, startDay: "samstag", weekData: wd},
		"februar":   MonthData{numberOfDays: 28, startDay: "dienstag", weekData: wd},
		"märz":      MonthData{numberOfDays: 31, startDay: "dienstag", weekData: wd},
		"april":     MonthData{numberOfDays: 30, startDay: "freitag", weekData: wd},
		"mai":       MonthData{numberOfDays: 31, startDay: "sonntag", weekData: wd},
		"juni":      MonthData{numberOfDays: 30, startDay: "mittwoch", weekData: wd},
		"juli":      MonthData{numberOfDays: 31, startDay: "freitag", weekData: wd},
		"august":    MonthData{numberOfDays: 31, startDay: "montag", weekData: wd},
		"september": MonthData{numberOfDays: 30, startDay: "donnerstag", weekData: wd},
		"oktober":   MonthData{numberOfDays: 31, startDay: "samstag", weekData: wd},
		"november":  MonthData{numberOfDays: 30, startDay: "dienstag", weekData: wd},
		"dezember":  MonthData{numberOfDays: 31, startDay: "donnerstag", weekData: wd},
	}

	month := flag.String("month", "default", "Monat")
	output := flag.String("output", "default", "Ausgabe")
	flag.Parse()

	myfile, err := os.Create(*output)
	if err != nil {
		fmt.Println(err)
	}
	defer myfile.Close()

	if *month != "" {
		monthTitle := "----" + *month + "----"
		monthOutput := data[*month].print()

		myfile.WriteString(monthTitle)
		myfile.WriteString(monthOutput)
		fmt.Println(monthTitle)
		fmt.Println(monthOutput)
	}
}
