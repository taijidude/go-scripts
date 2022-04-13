package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var created *os.File
	//delete a folder recursively
	err := os.RemoveAll("cptrgt")
	if err != nil {
		log.Fatal(err)
	}

	created, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	//write to text file
	_, err = created.WriteString("Hello World")
	if err != nil {
		log.Fatal(err)
	}

	//Close the file so it can be used later
	created.Close()

	//read from text file
	anotherFile, err := os.Open("anotherFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(anotherFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//create a new Folder
	if err := os.Mkdir("cptrgt", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	//copy a file
	bytesRead, err := ioutil.ReadFile("anotherFile.txt")
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile("cptrgt/anotherFile.txt", bytesRead, os.ModePerm)

	//Move a file
	err = os.Rename("test.txt", "cptrgt/test.txt")
	if err != nil {
		log.Fatal(err)
	}

	created, err = os.Create("toDelete.txt")
	if err != nil {
		log.Fatal(err)
	}
	created.Close()

	e := os.Remove("toDelete.txt")
	if e != nil {
		log.Fatal(e)
	}
}
