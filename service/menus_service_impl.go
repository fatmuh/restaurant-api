package service

import (
	"github.com/go-playground/validator/v10"
	"roastkuy-api/data/response"
	"roastkuy-api/repository"
)

type MenusServiceImpl struct {
	MenusRepository repository.MenusRepository
	Validate        *validator.Validate
}

func NewMenusServiceImpl(menusRepository repository.MenusRepository, validate *validator.Validate) MenusService {
	return &MenusServiceImpl{MenusRepository: menusRepository, Validate: validate}
}

func (m MenusServiceImpl) FindByOutletId(outletId int) ([]response.MenusResponse, error) {
	menuData, err := m.MenusRepository.FindByOutletId(outletId)
	if err != nil {
		return nil, err
	}

	var menuResponses []response.MenusResponse
	for _, menu := range menuData {
		menuResponse := response.MenusResponse{
			Id:          menu.Id,
			MenuName:    menu.MenuName,
			Description: menu.Description,
			Category:    menu.Category.NameCategory,
			Image:       menu.Image,
			Price:       menu.Price,
			Discount:    menu.Discount,
		}
		menuResponses = append(menuResponses, menuResponse)
	}

	return menuResponses, nil
}

func (m MenusServiceImpl) FindById(menuId int) (response.MenusResponse, error) {
	menuData, err := m.MenusRepository.FindById(menuId)
	if err != nil {
		return response.MenusResponse{}, err
	}

	menuResponse := response.MenusResponse{
		Id:          menuData.Id,
		MenuName:    menuData.MenuName,
		Description: menuData.Description,
		Category:    menuData.Category.NameCategory,
		Image:       menuData.Image,
		Price:       menuData.Price,
		Discount:    menuData.Discount,
	}

	return menuResponse, nil
}
