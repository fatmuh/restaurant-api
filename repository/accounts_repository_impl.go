package repository

import (
	"errors"
	"gorm.io/gorm"
	"roastkuy-api/data/request"
	"roastkuy-api/helper"
	"roastkuy-api/model"
)

type AccountsRepositoryImpl struct {
	Db *gorm.DB
}

func NewAccountsRepositoryImpl(Db *gorm.DB) AccountsRepository {
	return &AccountsRepositoryImpl{Db: Db}
}

func (a AccountsRepositoryImpl) Save(accounts model.Accounts) {
	result := a.Db.Create(&accounts)
	helper.ErrorPanic(result.Error)
}

func (a AccountsRepositoryImpl) Update(accounts model.Accounts) error {
	var updateAccount = request.UpdateAccountRequest{
		Id:    accounts.Id,
		Name:  accounts.Email,
		Phone: accounts.Phone,
	}

	result := a.Db.Model(&accounts).Updates(updateAccount)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (a AccountsRepositoryImpl) Delete(accountId int) error {
	var users model.Accounts
	result := a.Db.Where("id = ?", accountId).Delete(&users)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Account is not found")
	}
	return nil
}

func (a AccountsRepositoryImpl) FindById(accountId int) (model.Accounts, error) {
	var account model.Accounts
	result := a.Db.First(&account, accountId) // Use First instead of Find to ensure single record is fetched
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return account, errors.New("User is not found")
		}
		return account, result.Error
	}
	return account, nil
}

func (a AccountsRepositoryImpl) FindByEmail(email string) (model.Accounts, error) {
	var user model.Accounts
	result := a.Db.First(&user, "email = ?", email) // Use First instead of Find to ensure single record is fetched
	if result.Error != nil {
		return user, errors.New("Invalid email or password")
	}
	return user, nil
}
