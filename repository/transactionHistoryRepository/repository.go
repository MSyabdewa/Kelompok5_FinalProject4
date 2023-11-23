package transactionhistoryrepository

import (
	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
)

type Repository interface {
	CreateTransaction(transactionPayload *dto.NewTransactionHistoryRequest) (*dto.NewTransactionHistoryResponse, errs.Error)
	GetTransactionUser(userId int) (*[]dto.GetTransactionUser, errs.Error)
	GetTransactionAdmin() (*[]dto.GetTransactionAdmin, errs.Error)
}
