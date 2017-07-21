package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Person is a person
type Person struct {
	Name                               string
	Surname                            string
	Age                                int
	AVeryVeryLongData                  string
	AVeryVeryLongDataAVeryVeryLongData string
}

func r(w http.ResponseWriter, r *http.Request) {
	defer fmt.Println("Réponse envoyée !")
	p := Person{Name: "Deprez", Surname: "Adrien", Age: 23}
	payload, err := json.Marshal(p)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	w.Header().Add("Content-Type", "application/json")
	w.Write(payload)
}

func main() {
	http.HandleFunc("/", r)
	http.ListenAndServe(":8080", nil)
}
