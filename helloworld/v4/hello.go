package main

import (
	"fmt"
)

const german = "German"
const helloPrefix = "Hello, "
const germanHelloPrefix = "Hallo, "

// Hello takes a name and greets you
func Hello(n string, l string) string {

	if n == "" {
		n = "World"
	}

	if l == german {
		return germanHelloPrefix + n
	}

	return helloPrefix + n
}

func main() {
	fmt.Println(Hello("Tobi", ""))
}
