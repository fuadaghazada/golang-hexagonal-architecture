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

func (ah AccountHandler) NewTransaction(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	accountId := vars["accountId"]

	var newTransactionDto dto.NewTransactionDto
	if err := json.NewDecoder(request.Body).Decode(&newTransactionDto); err != nil {
		writeResponse(writer, http.StatusBadRequest, err.Error())
	} else {
		newTransactionDto.AccountId = accountId

		account, err := ah.service.NewTransaction(newTransactionDto)
		if err != nil {
			writeResponse(writer, err.Code, err.AsMessage())
		} else {
			writeResponse(writer, http.StatusCreated, account)
		}
	}
}