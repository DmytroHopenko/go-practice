package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")

		if name == " " || name == "" {
			name = "Guest"
		}

		fmt.Fprintf(w, "Hello, %s!", name)
	})

	http.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		current := time.Now().Format("2006-01-02 15:04:05")

		fmt.Fprintf(w, "Current time: %s", current)
	})

	fmt.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
