package animals

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Species int

var savedAnimals []Animal

const (
	Pig Species = iota
	Horse
	Cow
)

var SpeciesNames = [3]string{
	Pig:   "Pig",
	Horse: "Horse",
	Cow:   "Cow",
}

type Animal struct {
	Species Species
	Name    string
	Age     int
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.Encode(savedAnimals)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	animal := new(Animal)

	fmt.Sscanf(r.FormValue("species"), "%d", &animal.Species)
	fmt.Sscanf(r.FormValue("name"), "%s", &animal.Name)
	fmt.Sscanf(r.FormValue("age"), "%d", &animal.Age)

	savedAnimals = append(savedAnimals, *animal)

	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(animal)
}
