package dto

import (
	"time"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/entity"
)

type NewTransactionHistoryRequest struct {
	UserID    int
	ProductID int `json:"product_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required,gt=0"`
}

type NewTransactionHistoryResponse struct {
	TotalPrice   int    `json:"total_price"`
	Quantity     int    `json:"quantity"`
	ProductTitle string `json:"product_title"`
}

type TransactionHistoryBill struct {
	Message         string                        `json:"message"`
	TransactionBill NewTransactionHistoryResponse `json:"transaction_bill"`
}

type GetTransactionUser struct {
	ID         int `json:"id"`
	ProductID  int `json:"product_id"`
	UserID     int `json:"user_id"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
	Product    entity.Product
}

type GetTransactionAdmin struct {
	ID         int `json:"id"`
	ProductID  int `json:"product_id"`
	UserID     int `json:"user_id"`
	Quantity   int `json:"quantity"`
	TotalPrice int `json:"total_price"`
	Product    entity.Product
	User       TransactionUser
}

type TransactionUser struct {
	ID        int    `json:"id"`
	Email     string `json:"email"`
	FullName  string `json:"full_name"`
	Balance   int    `json:"balance"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
