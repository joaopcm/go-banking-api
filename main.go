package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"full_name"`
	City string `json:"city"`
	Zipcode string `json:"zip_code"`
}

func main() {
	http.HandleFunc("/great", greet)
	http.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{Name: "João", City: "São Paulo", Zipcode: "00000000"},
		{Name: "John Doe", City: "New York", Zipcode: "00000000"},
	}

	w.Header().Add("Content-Type", "application/json")

	json.NewEncoder(w).Encode(customers)
}
