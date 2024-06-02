package model

type Accounts struct {
	Id           int    `gorm:"type:int;primary_key"`
	Name         string `gorm:"type:varchar(255);not null"`
	Email        string `gorm:"type:varchar(255);unique;not null"`
	Password     string `gorm:"type:varchar(255);not null"`
	Phone        string `gorm:"type:varchar(255);unique;not null"`
	Roles        int    `gorm:"type:int;not null"`
	MemberNumber string `gorm:"type:varchar(255);not null"`
}
