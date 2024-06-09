package repository

import (
	"errors"
	"gorm.io/gorm"
	"roastkuy-api/model"
)

type MenusRepositoryImpl struct {
	Db *gorm.DB
}

func NewMenusRepositoryImpl(db *gorm.DB) MenusRepository {
	return &MenusRepositoryImpl{Db: db}
}

func (m MenusRepositoryImpl) BeginTransaction() *gorm.DB {
	return m.Db.Begin()
}

func (m MenusRepositoryImpl) FindByOutletId(outletId int) ([]model.Menus, error) {
	var menus []model.Menus
	result := m.Db.Preload("Category").Find(&menus, "outlet_id = ?", outletId) // Ensure single record is fetched by matching slug
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Outlet is not found")
		}
		return nil, result.Error
	}
	return menus, nil
}

func (m MenusRepositoryImpl) FindById(menuId int) (model.Menus, error) {
	var menu model.Menus
	result := m.Db.Preload("Category").First(&menu, menuId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return menu, errors.New("Menu is not found")
		}
		return menu, result.Error
	}
	return menu, nil
}
