package main

import (
	"fmt"
)

const helloPrefix = "Hello, "

// Hello greets you
func Hello(n string) string {
	return helloPrefix + n
}

func main() {
	fmt.Println(Hello("Tobi"))
}
