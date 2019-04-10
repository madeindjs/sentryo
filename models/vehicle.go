package models

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

func AllVehicles() Vehicles {
	rows, error := GetDatabase().Query("SELECT id, name, model FROM vehicles")

	vehicles := Vehicles{}

	if error != nil {
		panic(error)
	}

	for rows.Next() {
		vehicle := Vehicle{}

		rows.Scan(&vehicle.Id, &vehicle.Name, &vehicle.Model)
		vehicles = append(vehicles, vehicle)
	}

	return vehicles
}

func FindVehicle() Vehicle {
	vehicle := Vehicle{}
	// rows, _ := database.Query("SELECT name, model FROM vehicles")
	// var name string
	// var model string
	// for rows.Next() {
	// 	rows.Scan(&name, &model)
	// 	fmt.Println(": " + name + " " + model)
	// }
	return vehicle
}
