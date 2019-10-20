package main

import (
	"fmt"
)

func main() {

	planets := make(map[string]string)
	planets["Mars"] = "The Red one"
	planets["pluto"] = "The not real one"

	fmt.Println("welcome to space Quest!")
	fmt.Println("What is your name?")
	fmt.Println("John")
	fmt.Println("Hello John, do you want me to pick a planet for you?(Y/N)")

	var input string
	fmt.Scanln(&input)

	if input == "N" {
		fmt.Println("What planet do you want to go to?")
		fmt.Scanln(&input)
		for planet, info := range planets {
			if planet == input {
				fmt.Println(info)
			}
		}
	}
}
