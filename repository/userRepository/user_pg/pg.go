package user_pg

import (
	"database/sql"

	"github.com/MSyabdewa/Kelompok5_FinalProject4/dto"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/entity"
	"github.com/MSyabdewa/Kelompok5_FinalProject4/pkg/errs"
	userrepository "github.com/MSyabdewa/Kelompok5_FinalProject4/repository/userRepository"
)

const (
	createUsers = `
	INSERT into users (full_name, email,password) 
		VALUES ($1, $2, $3)
	RETURNING
		id,full_name,email,password,balance,created_at
	`

	login = `
	SELECT id, full_name, email, password, role, balance FROM users
	WHERE email = $1
	`

	topupBalance = `
	UPDATE users SET balance = balance + $1 WHERE id = $2
	RETURNING balance
	`

	countEmail = `
		SELECT COUNT(1) FROM users WHERE email = $1
	`

	getBalance = `
		SELECT balance FROM users WHERE id = $1
	`
)

type userPG struct {
	db *sql.DB
}

func NewUserPG(db *sql.DB) userrepository.Repository {
	return &userPG{
		db: db,
	}
}

func (userPG *userPG) CountEmail(email string) (int, errs.Error) {
	var count int

	err := userPG.db.QueryRow(countEmail, email).Scan(&count)
	if err != nil {
		return 0, errs.NewInternalServerError(err.Error())
	}

	return count, nil
}

func (userPG *userPG) CreateUser(u dto.NewUserRequest) (*dto.NewUserResponse, errs.Error) {
	// RETURNING id,full_name,email,password,balance,created_at
	var resp dto.NewUserResponse

	err := userPG.db.QueryRow(createUsers, u.FullName, u.Email, u.Password).Scan(
		&resp.ID, &resp.FullName, &resp.Email, &resp.Password,
		&resp.Balance, &resp.CreatedAt,
	)

	if err != nil {
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &resp, nil
}

func (userPG *userPG) Login(email string) (*entity.User, errs.Error) {
	var user entity.User

	err := userPG.db.QueryRow(login, email).Scan(
		&user.ID, &user.FullName, &user.Email, &user.Password, &user.Role, &user.Balance,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewUnauthenticatedError("Invalid Email or Password")
		}
		return nil, errs.NewInternalServerError(err.Error())
	}

	return &user, nil

}

func (userPG *userPG) TopUp(u dto.TopUpRequest) (int, errs.Error) {
	var balance int

	err := userPG.db.QueryRow(topupBalance, u.Balance, u.ID).Scan(
		&balance,
	)

	if err != nil {
		return 0, errs.NewInternalServerError(err.Error())
	}

	return balance, nil
}

func (userPG *userPG) GetBalance(id int) (int, errs.Error) {
	var balance int

	err := userPG.db.QueryRow(getBalance, id).Scan(
		&balance,
	)

	if err != nil {
		return 0, errs.NewInternalServerError(err.Error())
	}

	return balance, nil
}
