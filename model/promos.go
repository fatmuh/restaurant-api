package model

import "time"

type Promos struct {
	Id              int       `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	AccountId       int       `gorm:"type:int;foreign_key"`
	Image           string    `gorm:"type:varchar(255);"`
	Type            string    `gorm:"type:varchar(255);"`
	PromoName       string    `gorm:"type:varchar(255);"`
	Description     string    `gorm:"type:text;"`
	OutdatePromo    time.Time `gorm:"type:date"`
	DetailTutorial  string    `gorm:"type:text;"`
	DetailCondition string    `gorm:"type:text;"`
	CreatedAt       time.Time `gorm:"type:datetime;"`
	UpdatedAt       time.Time `gorm:"type:datetime;"`
	Outlet          Accounts  `gorm:"foreignkey:AccountId;association_foreignkey:Id"`
}

func (Promos) TableName() string {
	return "promo"
}
