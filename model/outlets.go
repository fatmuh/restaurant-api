package model

import "time"

type Outlets struct {
	Id                int       `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	Slug              string    `gorm:"type:varchar(255);unique_index"`
	OutletName        string    `gorm:"type:varchar(255);unique_index"`
	Address           string    `gorm:"type:text;unique_index"`
	Latitude          string    `gorm:"type:varchar(255);"`
	Longitude         string    `gorm:"type:varchar(255);"`
	OperationTime     string    `gorm:"type:varchar(255);"`
	Contact           string    `gorm:"type:varchar(255);"`
	GofoodLink        string    `gorm:"type:varchar(255);"`
	ShopeefoodLink    string    `gorm:"type:varchar(255);"`
	GrabfoodLink      string    `gorm:"type:varchar(255);"`
	TravelokaEatsLink string    `gorm:"type:varchar(255);"`
	CreatedAt         time.Time `gorm:"type:datetime;"`
	UpdatedAt         time.Time `gorm:"type:datetime;"`
}
