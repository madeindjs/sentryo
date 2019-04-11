package models

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

const databasePath = "swapi.dat"

var sqliteDatabase *sql.DB

func GetDatabase() *sql.DB {
	if sqliteDatabase == nil {

		database, error := sql.Open("sqlite3", databasePath)

		if error != nil {
			log.Fatal(error)
		}

		tableQueries := [...]string{
			"CREATE TABLE IF NOT EXISTS vehicles (name TEXT,model TEXT,manufacturer TEXT,cost_in_credits TEXT,length TEXT,max_atmosphering_speed TEXT,crew TEXT,passengers TEXT,cargo_capacity TEXT,consumables TEXT,vehicle_class TEXT,pilots TEXT,films TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
			"CREATE TABLE IF NOT EXISTS starships (name TEXT,model TEXT,manufacturer TEXT,cost_in_credits TEXT,length TEXT,max_atmosphering_speed TEXT,crew TEXT,passengers TEXT,cargo_capacity TEXT,consumables TEXT,hyperdrive_rating TEXT,MGLT TEXT,starship_class TEXT,pilots TEXT,films TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
			"CREATE TABLE IF NOT EXISTS people (name TEXT,height TEXT,mass TEXT,hair_color TEXT,skin_color TEXT,eye_color TEXT,birth_year TEXT,gender TEXT,homeworld TEXT,films TEXT,species TEXT,vehicles TEXT,starships TEXT,created TEXT,edited TEXT,url TEXT,id TEXT);",
		}

		for _, tableQuery := range tableQueries {
			statement, error := database.Prepare(tableQuery)
			if error != nil {
				log.Fatal(error)
			}

			statement.Exec()
		}

		sqliteDatabase = database // <--- NOT THREAD SAFE (see http://marcio.io/2015/07/singleton-pattern-in-go/)
	}
	return sqliteDatabase
}
