package models

import (
	"database/sql"
	"fmt"
)

type Starship struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Model string `json:"model"`

	// Omited for simplicity reasons :)
	// Manufacturer           string
	// Cost_in_credits        string
	// Length                 string
	// Max_atmosphering_speed string
	// Crew                   string
	// Passengers             string
	// Cargo_capacity         string
	// Consumables            string
	// Starship_class         string
	// Pilots                 string
	// Films                  string
	// Created                string
	// Edited                 string
	// Url                    string
	// MGLT                   string
	// Hyperdrive_rating      string
}

type Starships []Starship

/// Fetch all Starship from database
func AllStarships() Starships {
	rows, error := GetDatabase().Query("SELECT id, name, model FROM starships")

	starships := Starships{}

	if error != nil {
		panic(error)
	}

	for rows.Next() {
		starships = append(starships, createStarshipFromRow(rows))
	}

	rows.Close()

	return starships
}

/// Fetch a Starship from database
func FindStarship(id string) (Starship, error) {
	rows, error := GetDatabase().Query("SELECT id, name, model FROM starships WHERE id = ? LIMIT 1", id)
	defer rows.Close()

	if error != nil {
		return Starship{}, error
	}

	for rows.Next() {
		return createStarshipFromRow(rows), nil
	}

	return Starship{}, fmt.Errorf("Could not find a starship")
}

/// Initialize a Starship struct from a SQL response
func createStarshipFromRow(rows *sql.Rows) Starship {
	starship := Starship{}
	rows.Scan(&starship.Id, &starship.Name, &starship.Model)

	return starship
}
