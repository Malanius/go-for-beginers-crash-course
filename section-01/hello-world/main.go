package main

import "fmt"

func main() {
	saySomething("Hello, world!")
}

func saySomething(whatToSay string) {
	fmt.Println(whatToSay)
}
