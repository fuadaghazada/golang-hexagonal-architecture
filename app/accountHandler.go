package app

import (
	"encoding/json"
	"github.com/fuadaghazada/banking/dto"
	"github.com/fuadaghazada/banking/logger"
	"github.com/fuadaghazada/banking/service"
	"github.com/gorilla/mux"
	"net/http"
)

type AccountHandler struct {
	service service.AccountService
}

func (ah AccountHandler) NewAccount(writer http.ResponseWriter, request *http.Request) {
	var newAccountDto dto.NewAccountDto

	err := json.NewDecoder(request.Body).Decode(&newAccountDto)
	if err != nil {
		logger.Error("Invalid request payload")
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		newAccountDto.CustomerId = mux.Vars(request)["customerId"]
		account, appError := ah.service.NewAccount(newAccountDto)
		if appError != nil {
			writeResponse(writer, appError.Code, appError.AsMessage())
		} else {
			writeResponse(writer, http.StatusCreated, account)
		}
	}
}
