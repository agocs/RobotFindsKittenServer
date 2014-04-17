package main

import (
	"./response_generator"
	"fmt"
	"net/http"
)

func look(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", response_generator.GetResponse())
}

func main() {
	http.HandleFunc("/look/", look)
	http.ListenAndServe(":8080", nil)
}
