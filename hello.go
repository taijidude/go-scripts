package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	grade, err := strconv.Atoi(input)
	var result string
	if grade >= 60 {
		result = "passing"
	} else {
		result = "failing"
	}
	fmt.Println("grade: " + input)
	fmt.Println("status: " + result)
}
