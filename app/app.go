package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joaopcm/go-banking-api/domain"
	"github.com/joaopcm/go-banking-api/service"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
