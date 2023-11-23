package service

import (
	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/helpers"
	categoryrepository "github.com/MSyabdewa/Kelompok5_FinalProject4/repository/categoryRepository"
)

type categoryService struct {
	categoryRepo categoryrepository.Repository
}

type CategoryService interface {
	CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error)
	UpdateCategory(categoryId int, categoryPayload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.Error)
	GetCategories() (*dto.GetCategories, errs.Error)
	DeleteCategory(categoryId int) errs.Error
}

func NewCategoryService(categoryRepo categoryrepository.Repository) CategoryService {
	return &categoryService{categoryRepo: categoryRepo}
}

func (cs *categoryService) CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(categoryPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	res, err := cs.categoryRepo.CreateCategory(categoryPayload)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cs *categoryService) GetCategories() (*dto.GetCategories, errs.Error) {
	categories, err := cs.categoryRepo.GetCategories()
	if err != nil {
		return nil, err
	}
	return categories, nil
}

func (cs *categoryService) UpdateCategory(categoryId int, categoryPayload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.Error) {
	validateErr := helpers.ValidateStruct(categoryPayload)
	if validateErr != nil {
		return nil, validateErr
	}
	res, err := cs.categoryRepo.UpdateCategory(categoryId, categoryPayload)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (cs *categoryService) DeleteCategory(categoryId int) errs.Error {
	err := cs.categoryRepo.DeleteCategory(categoryId)
	if err != nil {
		return err
	}
	return nil
}
