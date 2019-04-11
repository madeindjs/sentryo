package models

import (
	"database/sql"
	"fmt"
	"github.com/golang/glog"
)

type Vehicle struct {
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
	// Vehicle_class          string
	// Pilots                 string
	// Films                  string
	// Created                string
	// Edited                 string
	// Url                    string
}

type Vehicles []Vehicle

/// Fetch all Vehicle from database
func AllVehicles() Vehicles {
	rows, error := GetDatabase().Query("SELECT id, name, model FROM vehicles")
	defer rows.Close()

	vehicles := Vehicles{}

	if error != nil {
		glog.Error(error)
	}

	for rows.Next() {
		vehicles = append(vehicles, CreateVehicleFromRow(rows))
	}

	return vehicles
}

/// Fetch a Vehicle from database
func FindVehicle(id string) (Vehicle, error) {
	rows, error := GetDatabase().Query("SELECT id, name, model FROM vehicles WHERE id = ? LIMIT 1", id)
	defer rows.Close()

	if error != nil {
		glog.Error(error)
		return Vehicle{}, error
	}

	for rows.Next() {
		return CreateVehicleFromRow(rows), nil
	}

	return Vehicle{}, fmt.Errorf("Could not find a vehicle")
}

/// Initialize a Vehicle struct from a SQL response
func CreateVehicleFromRow(rows *sql.Rows) Vehicle {
	vehicle := Vehicle{}
	rows.Scan(&vehicle.Id, &vehicle.Name, &vehicle.Model)

	return vehicle
}
