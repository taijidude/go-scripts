package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	layoutGerman = "18.08.1980"
)

func main() {
	entry := flag.String("t", "lol1234", "a string")
	flag.Parse()

	t1 := time.Now()

	fmt.Println(t1.Format("01.02.2006"))

	fmt.Println("--- Journal ---")
	fmt.Println(*entry)

}
