package service

import (
	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/helpers"
	transactionhistoryrepository "github.com/MSyabdewa/Kelompok5_FinalProject4/repository/transactionHistoryRepository"
)

type transactionService struct {
	transactionRepo transactionhistoryrepository.Repository
}

type TransactionService interface {
	CreateTransaction(transactionPayload *dto.NewTransactionHistoryRequest) (*dto.TransactionHistoryBill, errs.Error)
	GetTransactionUser(userId int) (*[]dto.GetTransactionUser, errs.Error)
	GetTransactionAdmin() (*[]dto.GetTransactionAdmin, errs.Error)
}

func NewTransactionService(transactionRepo transactionhistoryrepository.Repository) TransactionService {
	return &transactionService{transactionRepo: transactionRepo}
}

func (ts *transactionService) CreateTransaction(transactionPayload *dto.NewTransactionHistoryRequest) (*dto.TransactionHistoryBill, errs.Error) {
	validateErr := helpers.ValidateStruct(transactionPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	transaction, err := ts.transactionRepo.CreateTransaction(transactionPayload)
	if err != nil {
		return nil, err
	}
	res := dto.TransactionHistoryBill{
		Message:         "You have successfully purchased the product",
		TransactionBill: *transaction,
	}
	return &res, nil
}

func (ts *transactionService) GetTransactionUser(userId int) (*[]dto.GetTransactionUser, errs.Error) {
	transactions, err := ts.transactionRepo.GetTransactionUser(userId)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ts *transactionService) GetTransactionAdmin() (*[]dto.GetTransactionAdmin, errs.Error) {
	transactions, err := ts.transactionRepo.GetTransactionAdmin()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
