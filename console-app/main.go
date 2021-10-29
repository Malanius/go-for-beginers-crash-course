package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')

		userInput = strings.ReplaceAll(userInput, "\r\n", "")
		userInput = strings.ReplaceAll(userInput, "\n", "")

		if userInput == "quit" {
			break
		}

		fmt.Println(userInput)
	}
}
