package models

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
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

func (people *People) Insert() (sql.Result, error) {
	res, err := GetDatabase().Exec("INSERT INTO people(id, name, gender) values(?,?,?)", people.Id, people.Name, people.Gender)

	if err != nil {
		glog.Error(err)
		return res, err
	}

	return res, nil
}

func (people *People) Save() (sql.Result, error) {
	res, err := GetDatabase().Exec("UPDATE people SET name = ?, gender = ? WHERE id = ?", people.Name, people.Gender, people.Id)

	if err != nil {
		glog.Error(err)
		return res, err
	}

	return res, nil
}

func (people *People) Delete() (sql.Result, error) {
	res, err := GetDatabase().Exec("DELETE FROM people WHERE id = ?", people.Id)

	if err != nil {
		glog.Error(err)
		return res, err
	}

	return res, nil
}

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
	defer rows.Close()

	if error != nil {
		glog.Error(error)
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
