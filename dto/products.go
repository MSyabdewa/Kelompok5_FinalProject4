package dto

import "time"

type NewProductRequest struct {
	Title      string `json:"title" validate:"required"`
	Price      int    `json:"price" validate:"required,lt=50000000,gt=0"`
	Stock      int    `json:"stock" validate:"required,gt=5"`
	CategoryID int    `json:"category_id"`
}

type NewProductResponse struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetProduct struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
}

type GetProductResponse struct {
	Data []GetProduct `json:"data"`
}

type UpdateProductRequest struct {
	ID         int    `json:"id"`
	Title      string `json:"title" validate:"required"`
	Price      int    `json:"price" validate:"required,lt=50000000,gt=0"`
	Stock      int    `json:"stock" validate:"required,gt=5"`
	CategoryID int    `json:"category_id"`
}

type UpdatedProduct struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Price      int       `json:"price"`
	Stock      int       `json:"stock"`
	CategoryID int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated__at"`
}

type UpdateProductResponse struct {
	Product UpdatedProduct `json:"product"`
}

type DeleteProductsResponse struct {
	Message string `json:"message"`
}
