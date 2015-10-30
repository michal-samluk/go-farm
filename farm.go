package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Species int

const (
	Cattle Species = iota
	DomesticPig
	WildHorse
)

var speciesNames = [3]string{
	Cattle:      "Cattle",
	DomesticPig: "Domestic pig",
	WildHorse:   "Wild horse",
}

type Animal struct {
	Species Species
	Name    string
	Age     int
}

func requestAnimal() *Animal {
	animal := new(Animal)
	for index, name := range speciesNames {
		fmt.Printf("%d:%s\n", index, name)
	}
	fmt.Print("Select species: ")
	fmt.Scanf("%d", &animal.Species)

	fmt.Print("Name: ")
	fmt.Scanf("%s", &animal.Name)

	fmt.Print("Age: ")
	fmt.Scanf("%d", &animal.Age)
	return animal
}

func requestAnimals() (animals []*Animal) {
	for answer := "Yes"; answer == "Yes"; {
		animals = append(animals, requestAnimal())

		fmt.Print("Type 'Yes' if you want to continue: ")
		fmt.Scanf("%s", &answer)
	}
	return
}

func dumpAnimals(animals []*Animal) {
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
