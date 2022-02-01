package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	mux.HandleFunc("/great", greet)
	mux.HandleFunc("/customers", getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}