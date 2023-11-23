package productrepository

import (
	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
)

type Repository interface {
	CreateProduct(p dto.NewProductRequest) (*dto.NewProductResponse, errs.Error)
	UpdateProduct(p dto.UpdateProductRequest) (*dto.UpdateProductResponse, errs.Error)
	GetProducts() (*dto.GetProductResponse, errs.Error)
	DeleteProducts(id int) errs.Error
	CategoryIdExist(id int) (bool, errs.Error)
}
