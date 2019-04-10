package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

const port = ":8000"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello Sentryo")
	log.Print("GET /")
}

func vehiculesIndex(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	log.Print("GET /vehicules")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	mux.HandleFunc("/vehicules/", vehiculesIndex)

	log.Print("Start server on http://localhost", port)
	http.ListenAndServe(port, mux)
}
