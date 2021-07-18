package dto

type NewTransactionResponseDto struct {
	TransactionId string  `json:"transaction_id"`
	Balance       float64 `json:"balance"`
}
