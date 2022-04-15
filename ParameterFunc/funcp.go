package main

import "fmt"

func hello(first, second string) string {
	return first + second
}

func print(f func(string, string) string, first, second string) {
	fmt.Println(f(first, second))
}

func main() {
	print(hello, "Hello", "world")
}
