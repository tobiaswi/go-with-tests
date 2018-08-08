package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

// Hello takes a name and greets you
func Hello(n string) string {
	if n == "" {
		n = "World"
	}
	return helloPrefix + n
}

func main() {
	fmt.Println(Hello("Tobi"))
}
