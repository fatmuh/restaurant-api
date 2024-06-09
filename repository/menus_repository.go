package repository

import (
	"gorm.io/gorm"
	"roastkuy-api/model"
)

type MenusRepository interface {
	BeginTransaction() *gorm.DB
	FindByOutletId(outletId int) ([]model.Menus, error)
	FindById(menuId int) (model.Menus, error)
}
