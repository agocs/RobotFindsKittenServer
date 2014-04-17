package main

import (
	"fmt"
	"net/http"
)

func look(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", GetResponse())
}

func main() {
	http.HandleFunc("/look/", look)
	http.ListenAndServe(":8080", nil)
}
