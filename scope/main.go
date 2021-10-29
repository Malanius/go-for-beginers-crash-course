package main

import (
	"fmt"
	"scope/packageone"
)

var one = "One"

func main() {
	var one = "this is block level variable"
	fmt.Println(one)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)
	packageone.Exported()

	// No access as lowercase named things are private (seriously, wtf?)
	// secondString := packageone.privateVar
	// packageone.notExported()
}

func myFunc() {
	// var one = "the number one"
	fmt.Println(one)
}
