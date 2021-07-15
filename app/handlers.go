package app

import (
	"encoding/json"
	"encoding/xml"
	"github.com/fuadaghazada/banking/service"
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