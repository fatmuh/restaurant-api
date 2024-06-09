package service

import "roastkuy-api/data/response"

type MenusService interface {
	FindByOutletId(outletId int) ([]response.MenusResponse, error)
	FindById(menuId int) (response.MenusResponse, error)
}
