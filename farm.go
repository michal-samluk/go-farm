package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Species int

const (
	Cattle Species = iota
	DomesticPig
	WildHorse
)

var speciesNames = []string{
	Cattle:      "Cattle",
	DomesticPig: "Domestic pig",
	WildHorse:   "Wild horse",
}

type Animal struct {
	Species Species
	Name    string
	Age     int
}

func requestAnimal() (animal Animal) {
	fmt.Println("Select species:")
	for index, name := range speciesNames {
		fmt.Println(strconv.Itoa(index) + ": " + name)
	}

	var species Species
	fmt.Scanln(&species)
	animal.Species = species

	var name string
	fmt.Print("Name: ")
	fmt.Scanln(&name)
	animal.Name = name

	var age int
	fmt.Print("Age: ")
	fmt.Scanln(&age)
	animal.Age = age
	return
}

func requestAnimals() (animals []Animal) {
	for {
		animal := requestAnimal()
		animals = append(animals, animal)

		var answer string
		fmt.Print("Type 'Yes' if you want to continue: ")
		if fmt.Scanln(&answer); answer != "Yes" {
			break
		}
	}
	return
}

func dumpAnimals(animals []Animal) {
	file, err := os.Create("farm.json")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(animals)
}

func main() {
	dumpAnimals(requestAnimals())
}
