package animals

import (
	"encoding/json"
	"fmt"
	"github.com/michal-samluk/farm/app"
	"net/http"
)

type Species int

const (
	Pig Species = iota
	Horse
	Cow
)

type Animal struct {
	ID      int
	Species Species
	Name    string
	Age     int
}

func IndexHandler(a *app.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	rows, err := a.DB.Query("SELECT id, name, age, species FROM animals")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	var savedAnimals []Animal

	defer rows.Close()
	for rows.Next() {
		var a = Animal{}
		rows.Scan(&a.ID, &a.Name, &a.Age, &a.Species)
		savedAnimals = append(savedAnimals, a)
	}

	encoder := json.NewEncoder(w)
	encoder.Encode(savedAnimals)
	return http.StatusOK, nil
}

func CreateHandler(a *app.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	animal := new(Animal)

	fmt.Sscanf(r.FormValue("species"), "%d", &animal.Species)
	fmt.Sscanf(r.FormValue("name"), "%s", &animal.Name)
	fmt.Sscanf(r.FormValue("age"), "%d", &animal.Age)

	err := a.DB.QueryRow(`
    INSERT INTO animals(name, species, age)
    VALUES($1, $2, $3) RETURNING id`,
		animal.Name, animal.Species, animal.Age).Scan(&animal.ID)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	encoder := json.NewEncoder(w)
	if err = encoder.Encode(animal); err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}
