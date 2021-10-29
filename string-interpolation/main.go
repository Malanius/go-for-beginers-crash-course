package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	userName := readString("That is your name?")
	age := readInt("How old are you?")
	fmt.Println("Your name is " + userName + ". And you are", age, "years old.")
}

func promt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		promt()
		userinput, _ := reader.ReadString('\n')
		userinput = strings.ReplaceAll(userinput, "\r\n", "")
		userinput = strings.ReplaceAll(userinput, "\n", "")

		if userinput == "" {
			fmt.Print("Please enter a value!")
		} else {
			return userinput
		}
	}
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		promt()
		userinput, _ := reader.ReadString('\n')
		userinput = strings.ReplaceAll(userinput, "\r\n", "")
		userinput = strings.ReplaceAll(userinput, "\n", "")
		num, err := strconv.Atoi(userinput)
		if err != nil {
			fmt.Println("Please enter a whole number!")
		} else {
			return num
		}
	}
}
