package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

type Vehicule struct {
	Id    int    `json:"id"`
	Title string `json:"Title"`
}

type Vehicules []Vehicule

const port = ":8000"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Sentryo")
	log.Print("GET /")
}

func vehiculesIndex(w http.ResponseWriter, r *http.Request) {
	articles := Vehicules{
		Vehicule{Id: 1, Title: "Hello"},
		Vehicule{Id: 2, Title: "Hello 2"},
	}
	log.Print("GET /vehicules")

	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func vehiculesShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// handle error
		log.Fatal(err)
		os.Exit(2)
	}

	article := Vehicule{
		Id:    1,
		Title: "Hello",
	}

	log.Print("GET /vehicules/", id)

	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(article)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", hello)
	myRouter.HandleFunc("/vehicules/", vehiculesIndex)
	myRouter.HandleFunc("/vehicules/{id}", vehiculesShow)

	log.Print("Start server on http://localhost", port)
	log.Fatal(http.ListenAndServe(port, myRouter))

}

func main() {
	handleRequests()
}
