package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"gorm.io/gorm"
	config2 "roastkuy-api/config"
	"roastkuy-api/data/request"
	"roastkuy-api/helper"
	"roastkuy-api/model"
	"roastkuy-api/repository"
	"roastkuy-api/utils"
	"time"
)

type AuthServiceImpl struct {
	AccountsRepository repository.AccountsRepository
	Validate           *validator.Validate
	DB                 *gorm.DB
}

func NewAuthenticationServiceImpl(accountRepository repository.AccountsRepository, validate *validator.Validate, db *gorm.DB) AuthService {
	return &AuthServiceImpl{
		AccountsRepository: accountRepository,
		Validate:           validate,
		DB:                 db,
	}
}

func (a AuthServiceImpl) Login(account request.LoginRequest) (string, time.Time, error) {
	new_user, user_err := a.AccountsRepository.FindByEmail(account.Email)
	if user_err != nil {
		return "", time.Time{}, errors.New("ACCOUNT_NOT_FOUND")
	}

	config, _ := config2.LoadConfig(".")

	verify_error := utils.VerifyPassword(new_user.Password, account.Password)
	if verify_error != nil {
		return "", time.Time{}, errors.New("WRONG_PASSWORD")
	}

	token, exp, err_token := utils.GenerateToken(config.TokenExpiresIn, new_user.Id, config.TokenSecret)
	helper.ErrorPanic(err_token)
	return token, exp, nil
}

func (a AuthServiceImpl) Register(account request.CreateAccountRequest) {
	tx := a.AccountsRepository.BeginTransaction()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	hashedPassword, err := utils.HashPassword(account.Password)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	newuser := model.Accounts{
		Name:         account.Name,
		Email:        account.Email,
		Password:     hashedPassword,
		MemberNumber: utils.GenerateMemberNumber(a.DB),
		Phone:        account.Phone,
		Roles:        1,
		Uuid:         uuid.New().String(),
	}

	err = a.AccountsRepository.Save(tx, newuser)
	if err != nil {
		tx.Rollback()
		panic(err)
	}

	tx.Commit()
}
