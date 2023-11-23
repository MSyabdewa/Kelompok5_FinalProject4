package dto

import "time"

type NewCategoryRequest struct {
	Type string `json:"type" validate:"required"`
}

type NewCategoryResponse struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	CreatedAt         time.Time `json:"created_at"`
}

type UpdateCategoryResponse struct {
	ID                int       `json:"id"`
	Type              string    `json:"type"`
	SoldProductAmount int       `json:"sold_product_amount"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type GetCategories struct {
	Data []GetCategoriesWithProducts `json:"data"`
}

type GetCategoriesWithProducts struct {
	ID                int                   `json:"id"`
	Type              string                `json:"type"`
	SoldProductAmount int                   `json:"sold_product_amount"`
	CreatedAt         time.Time             `json:"created_at"`
	UpdatedAt         time.Time             `json:"updated_at"`
	Products          []GetProductsCategory `json:"Products"`
}

type GetProductsCategory struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
