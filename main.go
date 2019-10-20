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
	fmt.Println("we are in the solar system and it has 8 planets!")
	fmt.Println("What is your name?")
	fmt.Scanln(&input)
	goToPlanet("Hello "+input+", do you want me to pick a planet for you?(Y/N)", planets)

}

func tellabout(planet string, info string) {
	fmt.Println("Let me tell you about " + planet)
	fmt.Println(info)

}

func goToPlanet(intro string, planets map[int][]string) {
	fmt.Println(intro)
	var yn string
	fmt.Scanln(&yn)
	if yn == "N" || yn == "n" {
		fmt.Println("What planet do you want to go to?")
		fmt.Scanln(&yn)
		for _, list := range planets {
			if list[0] == yn {
				tellabout(list[0], list[1])
			}
		}
	} else if yn == "Y" || yn == "y" {
		randIndex := rand.Intn(2)
		tellabout(planets[randIndex][0], planets[randIndex][1])

	} else {
		goToPlanet("Invalid input, enter (Y/N)", planets)
	}
}
