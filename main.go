package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
)

type PlanetarySystem struct {
	Name    string   `json:"name"`
	Planets []Planet `json:"planets"`
}

type Planet struct {
	Name string `json:"name"`
	Info string `json:"description"`
}

func main() {

	file, err := os.Open("planetarySystem.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()

	byteVaue, _ := ioutil.ReadAll(file)

	var planetSyst PlanetarySystem

	json.Unmarshal(byteVaue, &planetSyst)

	var input string

	fmt.Println("welcome to space Quest!")
	fmt.Println("we are in the " + planetSyst.Name + " planetary system, it has " + strconv.Itoa(len(planetSyst.Planets)) + " planets!")
	fmt.Println("What is your name?")
	fmt.Scanln(&input)
	goToPlanet("Hello "+input+", do you want me to pick a planet for you?(Y/N)", planetSyst)

}

func tellabout(planet string, info string) {
	fmt.Println("Let me tell you about " + planet)
	fmt.Println(info)

}

func goToPlanet(intro string, planetSyst PlanetarySystem) {
	fmt.Println(intro)
	var input string
	fmt.Scanln(&input)
	if input == "N" || input == "n" {
		fmt.Println("What planet do you want to go to?")
		fmt.Scanln(&input)
		// var planet Planet
		for _, planet := range planetSyst.Planets {
			if planet.Name == input {
				tellabout(planet.Name, planet.Info)
			}
		}
	} else if input == "Y" || input == "y" {
		planet := planetSyst.Planets[rand.Intn(len(planetSyst.Planets))]
		tellabout(planet.Name, planet.Info)

	} else {
		goToPlanet("Invalid input, enter (Y/N)", planetSyst)
	}
}
