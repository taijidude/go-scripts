package main

import (
	"log"
	"os"
)

func handleError(err *error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := os.Rename("test.txt", "cptrgt/test.txt")
	handleError(&err)
	//instead of: 
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
