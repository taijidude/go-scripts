package main

import (
	"fmt"
)

func main() {
	myInt := 4
	myIntPointer := &myInt
	fmt.Println(myIntPointer)
	fmt.Println(*myIntPointer)
	myFloat := 98.6
	myFloatPointer := &myFloat
	fmt.Println(myFloatPointer)
	fmt.Println(*myFloatPointer)
	myBool := true
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)
	fmt.Println(*myBoolPointer)

	fmt.Println("----------------")
	updatePointer()
	fmt.Println("----------------")
	var myFloatPointer2 *float64 = createPointer()
	fmt.Println(*myFloatPointer2)
	fmt.Println("----------------")
	amount := 6
	double(&amount)
	fmt.Println(amount)
	fmt.Println("----------------")
}

func updatePointer() {
	myInt := 4
	fmt.Println(myInt)
	myIntPointer := &myInt
	*myIntPointer = 8
	fmt.Println(*myIntPointer)
	fmt.Println(myInt)
}

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func double(number *int) {
	*number *= 2
}
