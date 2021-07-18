package service

import (
	"github.com/fuadaghazada/banking/domain"
	"github.com/fuadaghazada/banking/dto"
	"github.com/fuadaghazada/banking/errs"
	"time"
)

const DATETIMELAYOUT = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.NewAccountDto) (*dto.NewAccountResponseDto, *errs.AppError)
	NewTransaction(dto.NewTransactionDto) (*dto.NewTransactionResponseDto, *errs.AppError)
}

type DefaultAccountService struct {
	accountRepo domain.AccountRepository
	transactionRepo domain.TransactionRepository
}

func (s DefaultAccountService) NewAccount(req dto.NewAccountDto) (*dto.NewAccountResponseDto, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	account := domain.Account{
		AccountId:   "",
		CustomerId:  req.CustomerId,
		OpeningDate: time.Now().Format(DATETIMELAYOUT),
		AccountType: req.AccountType,
		Amount:      req.Amount,
		Status:      "1",
	}

	createdAccount, err := s.accountRepo.Save(account)
	if err != nil {
		return nil, err
	}

	accountResponseDto := createdAccount.ToResponseDto()

	return &accountResponseDto, nil
}

func (s DefaultAccountService) NewTransaction(req dto.NewTransactionDto) (*dto.NewTransactionResponseDto, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	if req.IsTypeWithdrawal() {
		account, err := s.accountRepo.FindById(req.AccountId)
		if err != nil {
			return nil, err
		}
		if !account.CanWithdraw(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance for withdraw")
		}
	}

	transaction := domain.Transaction{
		AccountId: req.AccountId,
		Amount: req.Amount,
		TransactionDate: time.Now().Format(DATETIMELAYOUT),
		TransactionType: req.TransactionType,
	}

	newTransaction, err := s.transactionRepo.Save(transaction)
	if err != nil {
		return nil, err
	}

	account, err := s.accountRepo.FindById(newTransaction.AccountId)
	if err != nil {
		return nil, err
	}

	transactionDto := newTransaction.ToResponseDto(account.Amount)

	return &transactionDto, nil
}

func NewAccountService(accountRepo domain.AccountRepository, transactionRepo domain.TransactionRepository) DefaultAccountService {
	return DefaultAccountService{
		accountRepo: accountRepo,
		transactionRepo: transactionRepo,
	}
}