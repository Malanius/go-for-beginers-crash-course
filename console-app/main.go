package main

import (
	"fmt"
	"log"
	"strconv"

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

	coffes := make(map[int]string)
	coffes[1] = "Cappucino"
	coffes[2] = "Latte"
	coffes[3] = "Ameriacno"
	coffes[4] = "Mocha"
	coffes[5] = "Latte Macchiato"
	coffes[6] = "Espresso"

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

		if char == 'q' || char == 'Q' {
			break
		}

		i, _ := strconv.Atoi(string(char))
		text := fmt.Sprintf("You chose %s", coffes[i])
		fmt.Println(text)
	}

	fmt.Println("Program exitting...")
}
