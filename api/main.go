package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You've requested: %s\n", r.URL.Path)
	})

	http.ListenAndServe(":5000", nil)
}
