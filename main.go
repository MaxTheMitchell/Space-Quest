package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var input string
	planets := make(map[int][]string)
	planets[0] = append(planets[0], "Mars")
	planets[0] = append(planets[0], "The red one")
	planets[1] = append(planets[1], "pluto")
	planets[1] = append(planets[1], "the fake one")

	fmt.Println("welcome to space Quest!")
	fmt.Println("What is your name?")
	fmt.Scanln(&input)
	fmt.Println("Hello " + input + ", do you want me to pick a planet for you?(Y/N)")
	fmt.Scanln(&input)

	if input == "N" {
		fmt.Println("What planet do you want to go to?")
		fmt.Scanln(&input)
		for _, list := range planets {
			if list[0] == input {
				tellabout(list[0], list[1])
			}
		}
	} else {
		randIndex := rand.Intn(2)
		tellabout(planets[randIndex][0], planets[randIndex][1])

	}
}

func tellabout(planet string, info string) {
	fmt.Println("Let me tell you about " + planet)
	fmt.Println(info)

}
