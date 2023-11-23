package transaction_history_pg

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/entity"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
	transactionhistoryrepository "github.com/MSyabdewa/Kelompok5_FinalProject4/repository/transactionHistoryRepository"
)

type transactionPG struct {
	db *sql.DB
}

const (
	findUserByIdQuery = `
	SELECT * FROM users WHERE id = $1
	`

	findProductByIdQuery = `
	SELECT * FROM products WHERE id = $1
	`

	createTransactionQuery = `
	INSERT INTO transaction_histories (product_id, user_id, quantity, total_price)
	VALUES ($1, $2, $3, $4)
	`

	updateUserBalanceQuery = `
	UPDATE users SET balance = balance - $1 WHERE id = $2
	`

	updateProductQuery = `
	UPDATE products SET stock = stock - $1 WHERE id = $2
	`

	updateCategoryQuery = `
	UPDATE categories set sold_product_amount = sold_product_amount + $1 WHERE id = $2
	`

	getTransactionUserQuery = `
	SELECT 
	transaction_histories.id, product_id, user_id, quantity, total_price, 
	products.id, products.title, products.price, products.stock, products.category_id, products.created_at, products.updated_at
	FROM transaction_histories
	LEFT JOIN products ON transaction_histories.product_id = products.id
	WHERE transaction_histories.user_id = $1
	`

	getTransactionAdminQuery = `
	SELECT
	transaction_histories.id, product_id, user_id, quantity, total_price, 
	products.id, products.title, products.price, products.stock, products.category_id, products.created_at, products.updated_at, 
	users.id, users.email, users.full_name, users.balance, users.created_at, users.updated_at
	FROM transaction_histories
	LEFT JOIN products ON transaction_histories.product_id = products.id
	LEFT JOIN users ON transaction_histories.user_id = users.id;
	`
)

func NewTransactionPG(db *sql.DB) transactionhistoryrepository.Repository {
	return &transactionPG{db: db}
}

func (transactionPG *transactionPG) CreateTransaction(transactionPayload *dto.NewTransactionHistoryRequest) (*dto.NewTransactionHistoryResponse, errs.Error) {
	var user entity.User
	var product entity.Product
	err := transactionPG.db.QueryRow(findUserByIdQuery, transactionPayload.UserID).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Password, &user.Role, &user.Balance, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("user not found")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}
	err = transactionPG.db.QueryRow(findProductByIdQuery, transactionPayload.ProductID).Scan(
		&product.ID, &product.Title, &product.Price, &product.Stock, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("product with id: " + strconv.Itoa(transactionPayload.ProductID) + " not found")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	//Logic Check
	if product.Stock-transactionPayload.Quantity < 5 {
		return nil, errs.NewBadRequest("quantity requested exceed product stock")
	}
	totalPrice := transactionPayload.Quantity * product.Price
	if user.Balance < totalPrice {
		return nil, errs.NewBadRequest("insufficient user balance, total price: " + strconv.Itoa(totalPrice))
	}

	// Run Query
	ctx := context.Background()
	tx, _ := transactionPG.db.BeginTx(ctx, nil)

	_, err = tx.ExecContext(ctx, updateUserBalanceQuery, totalPrice, user.ID)
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	_, err = tx.ExecContext(ctx, updateProductQuery, transactionPayload.Quantity, product.ID)
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	_, err = tx.ExecContext(ctx, updateCategoryQuery, transactionPayload.Quantity, product.CategoryID)
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	_, err = tx.ExecContext(ctx, createTransactionQuery, product.ID, user.ID, transactionPayload.Quantity, totalPrice)
	if err != nil {
		tx.Rollback()
		return nil, errs.NewInternalServerError(err.Error())
	}

	err = tx.Commit()
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	res := dto.NewTransactionHistoryResponse{
		TotalPrice:   totalPrice,
		Quantity:     transactionPayload.Quantity,
		ProductTitle: product.Title,
	}
	return &res, nil
}

func (transactionPG *transactionPG) GetTransactionUser(userId int) (*[]dto.GetTransactionUser, errs.Error) {
	var res []dto.GetTransactionUser

	transactions, err := transactionPG.db.Query(getTransactionUserQuery, userId)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer transactions.Close()

	for transactions.Next() {
		var transaction dto.GetTransactionUser
		var product entity.Product
		err := transactions.Scan(
			&transaction.ID, &transaction.ProductID, &transaction.UserID, &transaction.Quantity, &transaction.TotalPrice,
			&product.ID, &product.Title, &product.Price, &product.Stock, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		transaction.Product = product
		res = append(res, transaction)
	}
	return &res, nil
}

func (transactionPG *transactionPG) GetTransactionAdmin() (*[]dto.GetTransactionAdmin, errs.Error) {
	var res []dto.GetTransactionAdmin

	transactions, err := transactionPG.db.Query(getTransactionAdminQuery)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer transactions.Close()

	for transactions.Next() {
		var transaction dto.GetTransactionAdmin
		var product entity.Product
		var user dto.TransactionUser
		err := transactions.Scan(
			&transaction.ID, &transaction.ProductID, &transaction.UserID, &transaction.Quantity, &transaction.TotalPrice,
			&product.ID, &product.Title, &product.Price, &product.Stock, &product.CategoryID, &product.CreatedAt, &product.UpdatedAt,
			&user.ID, &user.Email, &user.FullName, &user.Balance, &user.CreatedAt, &user.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		transaction.Product = product
		transaction.User = user
		res = append(res, transaction)
	}
	return &res, nil
}
