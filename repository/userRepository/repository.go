package userrepository

import (
	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/entity"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
)

type Repository interface {
	CreateUser(u dto.NewUserRequest) (*dto.NewUserResponse, errs.Error)
	Login(email string) (*entity.User, errs.Error)
	TopUp(u dto.TopUpRequest) (int, errs.Error)
	CountEmail(email string) (int, errs.Error)
	GetBalance(id int) (int, errs.Error)
}
