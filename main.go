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
	planetsyst := jsonToplaentsyst("planetarySystem.json")
	fmt.Println("welcome to space Quest!")
	fmt.Println("we are in the " + planetsyst.Name + " planetary system, it has " + strconv.Itoa(len(planetsyst.Planets)) + " planets!")
	fmt.Println("What is your name?")
	var username string
	fmt.Scanln(&username)
	goToPlanet("Hello "+username+", do you want me to pick a planet for you?(Y/N)", planetsyst)
}

func tellabout(planet string, info string) {
	fmt.Println("Let me tell you about " + planet)
	fmt.Println(info)

}

func goToPlanet(intro string, planetsyst PlanetarySystem) {
	fmt.Println(intro)
	var yn string
	fmt.Scanln(&yn)
	if yn == "N" || yn == "n" {
		enterplanet(planetsyst)
	} else if yn == "Y" || yn == "y" {
		planet := planetsyst.Planets[rand.Intn(len(planetsyst.Planets))]
		tellabout(planet.Name, planet.Info)

	} else {
		goToPlanet("Invalid input, enter (Y/N)", planetsyst)
	}
}

func enterplanet(planetsyst PlanetarySystem) {
	var userplanet string
	fmt.Println("What planet do you want to go to?")
	fmt.Scanln(&userplanet)
	for _, planet := range planetsyst.Planets {
		if planet.Name == userplanet {
			tellabout(planet.Name, planet.Info)
			return
		}
	}
	fmt.Println("invaild planet, enter another one!")
	enterplanet(planetsyst)
}

func jsonToplaentsyst(filename string) PlanetarySystem {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	byteVaue, _ := ioutil.ReadAll(file)
	var planetsyst PlanetarySystem
	json.Unmarshal(byteVaue, &planetsyst)
	return planetsyst
}
