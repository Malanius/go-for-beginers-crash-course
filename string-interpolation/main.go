package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	UserName       string
	Age            int
	FavoriteNumber float64
}

func main() {

	var user User
	user.UserName = readString("That is your name?")
	user.Age = readInt("How old are you?")
	user.FavoriteNumber = readFloat("What is your vaforite number?")

	fmt.Printf("Your name is %s. And you are %d years old.\nYour favorite number is %.2f", user.UserName, user.Age, user.FavoriteNumber)
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

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		promt()
		userinput, _ := reader.ReadString('\n')
		userinput = strings.ReplaceAll(userinput, "\r\n", "")
		userinput = strings.ReplaceAll(userinput, "\n", "")
		num, err := strconv.ParseFloat(userinput, 64)
		if err != nil {
			fmt.Println("Please enter a number!")
		} else {
			return num
		}
	}
}
