package main

import (
	"fmt"
	"net/http"
)

func look(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", GetResponse())
}

func main() {
	rand.Seed(time.Now().Unix())
	http.HandleFunc("/non-kitten-item/", look)
	http.ListenAndServe(":8080", nil)
}
