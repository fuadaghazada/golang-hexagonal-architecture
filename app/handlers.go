package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/fuadaghazada/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name" xml:"name"`
	City    string `json:"city" xml:"city"`
	ZipCode string `json:"zip_code" xml:"zip_code"`
}

type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(writer http.ResponseWriter, request *http.Request) {
	customers, _ := ch.service.GetAllCustomers()

	var err error = nil

	if request.Header.Get("Content-Type") == "application/xml" {
		writer.Header().Add("Content-Type", "application/xml")
		err = xml.NewEncoder(writer).Encode(customers)
	} else {
		writer.Header().Add("Content-Type", "application/json")
		err = json.NewEncoder(writer).Encode(customers)
	}

	if err != nil {
		return
	}
}

func (ch *CustomerHandler) getCustomer(writer http.ResponseWriter, request *http.Request) {
	customerId := mux.Vars(request)["customerId"]

	customer, err := ch.service.GetCustomerById(customerId)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, err.Error())
	} else {
		writer.Header().Add("Content-Type", "application/json")
		err := json.NewEncoder(writer).Encode(customer)
		if err != nil {
			return
		}
	}
}
