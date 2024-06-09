package service

import "roastkuy-api/data/response"

type PromosService interface {
	FindAll(accountID int) []response.PromosResponse
	FindRegular() []response.PromosResponse
}
