package model

import (
	"time"
)

// Menus structure represents the Menus table in the database
type Menus struct {
	Id          int        `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	OutletId    int        `gorm:"type:int;foreign_key"`
	CategoryId  int        `gorm:"type:int;foreign_key"`
	MenuName    string     `gorm:"type:varchar(255);"`
	Description string     `gorm:"type:text;"`
	Image       string     `gorm:"type:varchar(255);"`
	Price       int        `gorm:"type:int;"`
	Discount    float32    `gorm:"type:double"`
	CreatedAt   time.Time  `gorm:"type:datetime;"`
	UpdatedAt   time.Time  `gorm:"type:datetime;"`
	Outlet      Outlets    `gorm:"foreignkey:OutletId;association_foreignkey:Id"`
	Category    Categories `gorm:"foreignkey:CategoryId;association_foreignkey:Id"`
}

func (Menus) TableName() string {
	return "menu"
}
