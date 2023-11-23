package category_pg

import (
	"database/sql"
	"strconv"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
	categoryrepository "github.com/MSyabdewa/Kelompok5_FinalProject4/repository/categoryRepository"
)

const (
	createCategoryQuery = `
	INSERT INTO categories (type) VALUES ($1)
	RETURNING id, type, sold_product_amount, created_at
	`

	getCategoryQuery = `
	SELECT * FROM categories
	`
	getProductsByCategoryQuery = `
	SELECT * FROM products WHERE category_id = $1
	`

	updateCategoryQuery = `
	UPDATE categories SET type = $1 WHERE id = $2
	RETURNING id, type, sold_product_amount, updated_at
	`

	deleteCategoryQuery = `
	DELETE FROM categories WHERE id = $1
	`
)

type categoryPG struct {
	db *sql.DB
}

func NewCategoryPG(db *sql.DB) categoryrepository.Repository {
	return &categoryPG{db: db}
}

func (categoryPG *categoryPG) CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error) {
	var res dto.NewCategoryResponse
	err := categoryPG.db.QueryRow(createCategoryQuery, categoryPayload.Type).Scan(&res.ID, &res.Type, &res.SoldProductAmount, &res.CreatedAt)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &res, nil
}

func (categoryPG *categoryPG) GetCategories() (*dto.GetCategories, errs.Error) {
	var res dto.GetCategories
	var categories dto.GetCategoriesWithProducts
	var products dto.GetProductsCategory
	var temp int // for scan category_id

	categoryRows, err := categoryPG.db.Query(getCategoryQuery)
	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}
	defer categoryRows.Close()

	for categoryRows.Next() {
		err := categoryRows.Scan(
			&categories.ID, &categories.Type, &categories.SoldProductAmount, &categories.CreatedAt, &categories.UpdatedAt,
		)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		productRows, err := categoryPG.db.Query(getProductsByCategoryQuery, categories.ID)
		if err != nil {
			return nil, errs.NewInternalServerError(err.Error())
		}
		for productRows.Next() {
			err = productRows.Scan(&products.ID, &products.Title, &products.Price, &products.Stock, &temp, &products.CreatedAt, &products.UpdatedAt)
			if err != nil {
				return nil, errs.NewInternalServerError(err.Error())
			}
			categories.Products = append(categories.Products, products)
		}
		productRows.Close()
		res.Data = append(res.Data, categories)
		categories.Products = nil
	}
	return &res, nil
}

func (categoryPG *categoryPG) UpdateCategory(categoryId int, categoryPayload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.Error) {
	var res dto.UpdateCategoryResponse
	err := categoryPG.db.QueryRow(updateCategoryQuery, categoryPayload.Type, categoryId).Scan(
		&res.ID, &res.Type, &res.SoldProductAmount, &res.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Category with id: " + strconv.Itoa(categoryId) + " does not exist")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}
	return &res, nil
}

func (categoryPG *categoryPG) DeleteCategory(categoryId int) errs.Error {
	_, err := categoryPG.db.Exec(deleteCategoryQuery, categoryId)
	if err != nil {
		return errs.NewInternalServerError(err.Error())
	}
	return nil
}
