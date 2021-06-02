package main

//Hier stehen die Import Pfade der Packages
import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

//Ab der Main Methode werden die Namen der Packages verwendet
func main() {
	var target int
	success := false
	for x := 1; x <= 10; x++ {
		fmt.Println("Try " + strconv.Itoa(x) + "/10")
		seconds := time.Now().Unix()
		rand.Seed(seconds)
		target = rand.Intn(100)
		fmt.Println("I've chosen a random number between 1 and 100.")
		fmt.Println("Can you guess it?")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}

		if guess < target {
			fmt.Println("Ooops. Your guess was LOW")
		} else if guess > target {
			fmt.Println("Ooops. Your guess was HIGH")
		} else {
			fmt.Println("Your guess was right! You win!")
			success = true
			break
		}
		if x == 10 {
			fmt.Println("That was the last try!")
		}
	}
	if success == false {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
