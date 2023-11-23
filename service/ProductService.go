package service

import (
	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/helpers"
	productrepository "github.com/MSyabdewa/Kelompok5_FinalProject4/repository/productRepository"
)

type ProductService interface {
	CreateProduct(p dto.NewProductRequest) (*dto.NewProductResponse, errs.Error)
	GetProducts() (*dto.GetProductResponse, errs.Error)
	UpdateProduct(p dto.UpdateProductRequest) (*dto.UpdateProductResponse, errs.Error)
	DeleteProduct(id int) (*dto.DeleteProductsResponse, errs.Error)
}

type productService struct {
	productRepo productrepository.Repository
}

func NewProductService(productRepo productrepository.Repository) ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (ps *productService) CreateProduct(p dto.NewProductRequest) (*dto.NewProductResponse, errs.Error) {

	validateErr := helpers.ValidateStruct(p)
	if validateErr != nil {
		return nil, validateErr
	}

	categoryExist, err := ps.productRepo.CategoryIdExist(p.CategoryID)
	if err != nil {
		return nil, err
	}

	if !categoryExist {
		return nil, errs.NewBadRequest("Category Not Exist")
	}

	resp, err := ps.productRepo.CreateProduct(p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ps *productService) GetProducts() (*dto.GetProductResponse, errs.Error) {
	resp, err := ps.productRepo.GetProducts()
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ps *productService) UpdateProduct(p dto.UpdateProductRequest) (*dto.UpdateProductResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(p)
	if validateErr != nil {
		return nil, validateErr
	}

	categoryExist, err := ps.productRepo.CategoryIdExist(p.CategoryID)
	if err != nil {
		return nil, err
	}

	if !categoryExist {
		return nil, errs.NewBadRequest("Category Not Exist")
	}

	resp, err := ps.productRepo.UpdateProduct(p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ps *productService) DeleteProduct(id int) (*dto.DeleteProductsResponse, errs.Error) {
	var resp dto.DeleteProductsResponse

	err := ps.productRepo.DeleteProducts(id)
	if err != nil {
		return nil, err
	}

	resp.Message = "Product Has Been Successfully Deleted"
	return &resp, nil
}
