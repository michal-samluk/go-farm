package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Animal struct {
	Species []string
	Name    string
	Age     int
}

func (a *Animal) SetSpecies(species string) {
	speciesArr := strings.Split(species, ",")
	a.Species = speciesArr
}

func (a *Animal) SetAge(age string) {
	var err error
	if a.Age, err = strconv.Atoi(age); err != nil {
		panic(err)
	}
}

func (a *Animal) SetName(name string) {
	a.Name = name
}

func requestAnimal() (animal Animal) {
	animal.SetSpecies(getInput("Species (seperated with comma): "))
	animal.SetName(getInput("Name: "))
	animal.SetAge(getInput("Age: "))
	return
}

func requestAnimals() (animals []Animal) {
	for {
		animal := requestAnimal()
		animals = append(animals, animal)

		if getInput("Type 'Yes' if you want to continue: ") != "Yes" {
			break
		}
	}
	return
}

func getInput(message string) (input string) {
	fmt.Print(message)
	fmt.Scanln(&input)
	return
}

func dumpAnimals(animals []Animal) {
	file, err := os.Create("farm.json")

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(animals)
}

func main() {
	animals := requestAnimals()
	dumpAnimals(animals)
}
