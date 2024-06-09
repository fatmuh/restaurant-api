package repository

import (
	"gorm.io/gorm"
	"roastkuy-api/model"
)

type OutletsRepository interface {
	BeginTransaction() *gorm.DB
	FindAll() []model.Outlets
	FindBySlug(slug string) (*model.Outlets, error)
}
