package repository

import (
	"errors"
	"gorm.io/gorm"
	"roastkuy-api/helper"
	"roastkuy-api/model"
)

type OutletsRepositoryImpl struct {
	Db *gorm.DB
}

func NewOutletsRepositoryImpl(db *gorm.DB) OutletsRepository {
	return &OutletsRepositoryImpl{Db: db}
}

func (o OutletsRepositoryImpl) BeginTransaction() *gorm.DB {
	return o.Db.Begin()
}

func (o OutletsRepositoryImpl) FindAll() []model.Outlets {
	var outlets []model.Outlets
	result := o.Db.Find(&outlets)
	helper.ErrorPanic(result.Error)
	return outlets
}

func (o OutletsRepositoryImpl) FindBySlug(slug string) (*model.Outlets, error) {
	var outlet model.Outlets
	result := o.Db.First(&outlet, "slug = ?", slug) // Ensure single record is fetched by matching slug
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Slug is not found")
		}
		return nil, result.Error
	}
	return &outlet, nil
}
