package app

import (
	"github.com/fuadaghazada/banking/domain"
	"github.com/fuadaghazada/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	// Wiring
	//ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandler{service: service.NewCustomerService(domain.NewCustomerRepositoryDb())}

	// Routes
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		return
	}
}
