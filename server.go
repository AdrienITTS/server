package main

import (
	"encoding/json"
	"net/http"
)

// Person is a person
type Person struct {
	Name    string
	Surname string
	Age     int
}

func r(w http.ResponseWriter, r *http.Request) {
	p := Person{Name: "Deprez", Surname: "Adrien", Age: 23}
	payload, _ := json.Marshal(p)
	w.Header().Add("Content-Type", "application/json")
	w.Write(payload)
}

func main() {
	http.HandleFunc("/", r)
	http.ListenAndServe(":8080", nil)
}
