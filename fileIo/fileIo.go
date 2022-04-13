package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func readYNInput(prompt string) (bool, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	text, err := reader.ReadString('\n')
	if err != nil {
		return false, err
	}
	//Cleans up the Input from the cli
	text = strings.TrimSpace(text)
	if strings.ToLower(text) == "y" {
		return true, nil
	}
	return false, nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func main2() {
	var created *os.File
	//delete a folder recursively
	err := os.RemoveAll("cptrgt")
	if err != nil {
		log.Fatal(err)
	}

	fileName := "test.txt"
	//check if file does not exist
	//and ask the user if he wants to create it
	if !fileExists(fileName) {
		input, err := readYNInput(fileName + " does not exist! Should it be created?")
		if err != nil {
			log.Fatal(err)
		}

		if input {
			fmt.Println("Creating file")
		} else {
			fmt.Println("Exiting")
			os.Exit(1)
		}
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

	fileName = "anotherFile.txt"
	//Check if a file exists
	if fileExists(fileName) {
		input, err := readYNInput("The file " + fileName + " exists! Read the contents?")
		if err != nil {
			log.Fatal(err)
		}
		if input {
			fmt.Println("Reading file")
		} else {
			fmt.Println("Exiting")
			os.Exit(1)
		}
	}

	//read from text file
	anotherFile, err := os.Open(fileName)
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
}
