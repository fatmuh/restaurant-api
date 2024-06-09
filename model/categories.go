package model

import "time"

type Categories struct {
	Id           int       `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	NameCategory string    `gorm:"type:varchar(255);"`
	CreatedAt    time.Time `gorm:"type:datetime;"`
	UpdatedAt    time.Time `gorm:"type:datetime;"`
}

func (Categories) TableName() string {
	return "menu_category"
}
