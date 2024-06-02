package repository

import (
	"gorm.io/gorm"
	"roastkuy-api/model"
)

type AccountsRepository interface {
	BeginTransaction() *gorm.DB
	Save(tx *gorm.DB, accounts model.Accounts) error
	Update(accounts model.Accounts) error
	Delete(accountId int) error
	FindById(accountId int) (model.Accounts, error)
	FindByEmail(email string) (model.Accounts, error)
}
