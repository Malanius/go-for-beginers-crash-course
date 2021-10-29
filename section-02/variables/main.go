package main

import (
	"bufio"
	"fmt"
	"os"
)

// Constant
const promt = "and press ENTER when ready..."

func main() {
	// Declare and assign separately
	var firstNumber int
	firstNumber = 2

	// Declare and assign value
	var secondNumber = 5

	// declare and assign with := operator
	subtraction := 7

	// Without assigment, has default value, for int = 0
	var answer int

	reader:= bufio.NewReader(os.Stdin)

	fmt.Println("Gues the Number Game")
	fmt.Println("--------------------")
	fmt.Println()

	fmt.Println("Think of a number between 1 and 10", promt)
	reader.ReadString('\n')

	fmt.Println("Multiply your number by", firstNumber, promt)
	reader.ReadString('\n')

	fmt.Println("Now multipy the result by", secondNumber, promt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", promt)
	reader.ReadString('\n')

	fmt.Println("Now subtract", subtraction, promt)
	reader.ReadString('\n')

	answer = firstNumber * secondNumber - subtraction
	fmt.Println("The answer is", answer)
}
