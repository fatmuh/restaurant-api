package repository

import (
	"gorm.io/gorm"
	"roastkuy-api/helper"
	"roastkuy-api/model"
)

type PromosRepositoryImpl struct {
	Db *gorm.DB
}

func NewPromosRepositoryImpl(db *gorm.DB) PromosRepository {
	return &PromosRepositoryImpl{Db: db}
}

func (p PromosRepositoryImpl) FindAll(currentAccountID int) []model.Promos {
	var promos []model.Promos
	result := p.Db.Where("account_id = ? OR account_id = 0", currentAccountID).Find(&promos)
	helper.ErrorPanic(result.Error)
	return promos
}

func (p PromosRepositoryImpl) FindRegular() []model.Promos {
	var promos []model.Promos
	result := p.Db.Where("type = ?", "regular").Find(&promos)
	helper.ErrorPanic(result.Error)
	return promos
}
