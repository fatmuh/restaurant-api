package service

import "roastkuy-api/data/response"

type OutletsService interface {
	FindAll() []response.OutletsResponse
	FindBySlug(slug string) (response.OutletsResponse, error)
}
