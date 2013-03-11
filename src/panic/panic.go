package main

import (
	"os"
	"fmt"
)

var user = os.Getenv("USER")
func main() {
	test()
}

func test() {
	if user == "" {
		// panic("no value for $USER")
		fmt.Println("next statement")
	} else {
		panic("$USER is set")
		fmt.Println("Set USER")
	}
}