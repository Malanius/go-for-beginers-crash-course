package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	// defer runs after end of the function
	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("MENU")
	fmt.Println("----")
	fmt.Println("1 - Cappucino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Ameriacno")
	fmt.Println("4 - Mocha")
	fmt.Println("5 - Latte Macchiato")
	fmt.Println("6 - Espresso")
	fmt.Println("Q - Quit the program")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		text:= fmt.Sprintf("You chose %q", char)
		fmt.Println(text)

		if char == 'q' || char == 'Q' {
			break
		}

	}

	fmt.Println("Program exitting...")
}
