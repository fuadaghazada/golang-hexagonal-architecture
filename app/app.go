package app

import (
	"fmt"
	"github.com/fuadaghazada/banking/domain"
	"github.com/fuadaghazada/banking/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func sanityCheck() {
	if os.Getenv("SERVER_HOST") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Environment variables not found...")
	}
}

func Start() {
	sanityCheck()

	router := mux.NewRouter()

	// Wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerId:[0-9]+}", ch.getCustomer).Methods(http.MethodGet)

	host := os.Getenv("SERVER_HOST")
	port := os.Getenv("SERVER_PORT")
	serverAddr := fmt.Sprintf("%s:%s", host, port)
	err := http.ListenAndServe(serverAddr, router)
	if err != nil {
		return
	}
}
