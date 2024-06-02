package repository

import "roastkuy-api/model"

type AccountsRepository interface {
	Save(accounts model.Accounts)
	Update(accounts model.Accounts) error
	Delete(accountId int) error
	FindById(accountId int) (model.Accounts, error)
	FindByEmail(email string) (model.Accounts, error)
}
