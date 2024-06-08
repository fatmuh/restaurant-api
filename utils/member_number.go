package utils

import (
	"gorm.io/gorm"
	"roastkuy-api/model"
)

func GetLastMember(db *gorm.DB) *model.Accounts {
	var lastMember model.Accounts
	db.Order("id DESC").First(&lastMember)
	return &lastMember
}

func GenerateMemberNumber(db *gorm.DB) int {
	lastMember := GetLastMember(db)
	var lastGeneratedNumber int

	if lastMember != nil {
		lastGeneratedNumber = lastMember.MemberNumber
	} else {
		lastGeneratedNumber = 1000
	}

	lastGeneratedNumber++

	return lastGeneratedNumber
}
