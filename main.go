package main

import (
	"fmt"
	"io"
	"net/http"
)

const port = ":8000"

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)

	fmt.Printf("Start server on http://localhost%v\n", port)
	http.ListenAndServe(port, mux)
}
