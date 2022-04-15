package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter a grade:")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	grade, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade == 100 {
		status = "A+"
	} else if grade >= 60 {
		status = "Pass"
	} else {
		status = "Fail"
	}
	fmt.Println("A grade of", grade, "is a", status)

}
