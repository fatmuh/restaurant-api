package service

import (
	"errors"
	"github.com/go-playground/validator/v10"
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
}

func NewAuthenticationServiceImpl(accountRepository repository.AccountsRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		AccountsRepository: accountRepository,
		Validate:           validate,
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
	hashedPassword, err := utils.HashPassword(account.Password)
	if err != nil {
		panic(err)
	}

	newuser := model.Accounts{
		Name:     account.Name,
		Email:    account.Email,
		Password: hashedPassword,
		Phone:    account.Phone,
	}

	a.AccountsRepository.Save(newuser)
}
