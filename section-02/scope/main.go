package main

import "fmt"

var one = "One"

func main() {
	var one = "this is block level variable"
	fmt.Println(one)
	myFunc()
}

func myFunc() {
	// var one = "the number one"
	fmt.Println(one)
}
