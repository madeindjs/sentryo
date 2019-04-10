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

type Article struct {
	id    int    `json:"id"`
	title string `json:"title"`
}

type Articles []Article

const port = ":8000"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Sentryo")
	log.Print("GET /")
}

func vehiculesIndex(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{id: 1, title: "Hello"},
		Article{id: 2, title: "Hello 2"},
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

	article := Article{
		id:    1,
		title: "Hello",
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
