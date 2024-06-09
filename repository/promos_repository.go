package repository

import "roastkuy-api/model"

type PromosRepository interface {
	FindAll(accountID int) []model.Promos
	FindRegular() []model.Promos
}
