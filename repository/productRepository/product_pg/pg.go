package product_pg

import (
	"database/sql"

	productrepository "github.com/MSyabdewa/Kelompok5_FinalProject4/repository/productRepository"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
)

const (
	createProduct = `
		INSERT INTO products
			(title,price,stock,category_id)
		VALUES
			($1,$2,$3,$4)
		RETURNING
			id,title,price,stock,category_id,created_at
	`

	getProducts = `
		SELECT 
			id,title,price,stock,category_id,created_at
		FROM products
	`

	updateProduct = `
		UPDATE products SET
			title = $1,price = $2,
			stock = $3,category_id = $4,
			updated_at = current_timestamp
		WHERE id = $5
		RETURNING
			id,title,price,stock,category_id,created_at,updated_at
	`

	deleteProduct = `
		DELETE FROM products WHERE id = $1
	`

	countCategory = `
		SELECT COUNT(1) FROM categories WHERE id = $1
	`
)

type productPG struct {
	db *sql.DB
}

func NewProductPG(db *sql.DB) productrepository.Repository {
	return &productPG{
		db: db,
	}
}

func (productPG *productPG) CreateProduct(p dto.NewProductRequest) (*dto.NewProductResponse, errs.Error) {
	var resp dto.NewProductResponse

	err := productPG.db.QueryRow(createProduct, p.Title, p.Price, p.Stock, p.CategoryID).Scan(
		&resp.ID, &resp.Title, &resp.Price, &resp.Stock,
		&resp.CategoryID, &resp.CreatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &resp, nil
}

func (productPG *productPG) UpdateProduct(p dto.UpdateProductRequest) (*dto.UpdateProductResponse, errs.Error) {
	var product dto.UpdatedProduct
	var resp dto.UpdateProductResponse

	err := productPG.db.QueryRow(updateProduct, p.Title, p.Price, p.Stock, p.CategoryID, p.ID).Scan(
		&product.ID, &product.Title, &product.Price,
		&product.Stock, &product.CategoryID, &product.CreatedAt,
		&product.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Product With Such ID Is Not Exist")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	resp.Product = product

	return &resp, nil
}

func (productPG *productPG) GetProducts() (*dto.GetProductResponse, errs.Error) {
	var resp dto.GetProductResponse
	var product dto.GetProduct

	rows, err := productPG.db.Query(getProducts)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&product.ID, &product.Title, &product.Price,
			&product.Stock, &product.CategoryID, &product.CreatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}

		resp.Data = append(resp.Data, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &resp, nil
}

func (productPG *productPG) DeleteProducts(id int) errs.Error {
	_, err := productPG.db.Exec(deleteProduct, id)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}

	return nil
}

func (productPG *productPG) CategoryIdExist(id int) (bool, errs.Error) {
	var count int

	err := productPG.db.QueryRow(countCategory, id).Scan(
		&count,
	)

	if err != nil {

		return false, errs.NewInternalServerError(err.Error())
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
