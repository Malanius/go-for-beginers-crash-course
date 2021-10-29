package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Constant
const promt = "and don't type your number in, just press ENTER..."

func main() {
	// Seeding rng
	rand.Seed(time.Now().UnixNano())

	// Declare and assign separately
	var firstNumber int
	firstNumber = rand.Intn(8) + 2

	// Declare and assign value
	var secondNumber = rand.Intn(8) + 2

	// declare and assign with := operator
	subtraction := rand.Intn(8) + 2

	// Without assigment, has default value, for int = 0
	var answer int
	answer = firstNumber*secondNumber - subtraction

	playGame(firstNumber, secondNumber, subtraction, answer)
}

func playGame(firstNumber int, secondNumber int, subtraction int, answer int) {
	reader := bufio.NewReader(os.Stdin)

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

	fmt.Println("The answer is", answer)
}
