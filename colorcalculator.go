package main

import (
	"fmt"
	"log"
)

var metersPerLiter float64

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}

	area := width * height
	return area / metersPerLiter, nil
}

func main() {
	metersPerLiter = 10.0
	amount, err := paintNeeded(4.2, -3.0)
	if err != nil {
		log.Fatal(err) //<-- print error and exit the program
	} else {
		fmt.Printf("%.2f liters needed\n", amount)
	}
}
