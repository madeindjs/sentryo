package models

import (
	"database/sql"
	"fmt"
)

type People struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`

	// Omited for simplicity reasons :)
	// Height string
	// Mass string
	// Hair_color string
	// Skin_color string
	// Eye_color string
	// Birth_year string
	// Homeworld string
	// Films string
	// Species string
	// Vehicles string
	// Starships string
	// Created string
	// Edited string
	// Url string
}

type Peoples []People

/// Fetch all People from database
func AllPeoples() Peoples {
	rows, error := GetDatabase().Query("SELECT id, name, gender FROM people")

	peoples := Peoples{}

	if error != nil {
		panic(error)
	}

	for rows.Next() {
		peoples = append(peoples, createPeopleFromRow(rows))
	}

	return peoples
}

/// Fetch a People from database
func FindPeople(id string) (People, error) {
	rows, error := GetDatabase().Query("SELECT id, name, gender FROM people WHERE id = ? LIMIT 1", id)

	if error != nil {
		return People{}, error
	}

	for rows.Next() {
		return createPeopleFromRow(rows), nil
	}

	return People{}, fmt.Errorf("Could not find a people")
}

/// Initialize a People struct from a SQL response
func createPeopleFromRow(rows *sql.Rows) People {
	people := People{}
	rows.Scan(&people.Id, &people.Name, &people.Gender)

	return people
}
