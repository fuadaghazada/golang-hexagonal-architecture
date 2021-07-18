package app

import (
	"encoding/json"
	"github.com/fuadaghazada/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	status := request.URL.Query().Get("status")
	customers, err := ch.service.GetAllCustomers(status)

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customers)
	}
}

func (ch *CustomerHandler) getCustomer(writer http.ResponseWriter, request *http.Request) {
	customerId := mux.Vars(request)["customerId"]
	customer, err := ch.service.GetCustomerById(customerId)

	if err != nil {
		writeResponse(writer, err.Code, err.AsMessage())
	} else {
		writeResponse(writer, http.StatusOK, customer)
	}
}

func writeResponse(writer http.ResponseWriter, code int, data interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(code)
	err := json.NewEncoder(writer).Encode(data)
	if err != nil {
		panic(err)
	}
}
