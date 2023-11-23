package categoryrepository

import (
	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
)

type Repository interface {
	CreateCategory(categoryPayload *dto.NewCategoryRequest) (*dto.NewCategoryResponse, errs.Error)
	UpdateCategory(categoryId int, categoryPayload *dto.NewCategoryRequest) (*dto.UpdateCategoryResponse, errs.Error)
	GetCategories() (*dto.GetCategories, errs.Error)
	DeleteCategory(categoryId int) errs.Error
}
