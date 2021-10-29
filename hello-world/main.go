package main

import "fmt"

func main() {
	// var message string
	// message = "Hello, world!"
	message := "Hello, world!"
	saySomething(message)
}

func saySomething(whatToSay string) {
	fmt.Println(whatToSay)
}
