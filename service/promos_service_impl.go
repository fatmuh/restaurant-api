package service

import (
	"github.com/go-playground/validator/v10"
	"roastkuy-api/data/response"
	"roastkuy-api/repository"
)

type PromosServiceImpl struct {
	PromosRepository repository.PromosRepository
	Validate         *validator.Validate
}

func NewPromosServiceImpl(promosRepository repository.PromosRepository, validate *validator.Validate) PromosService {
	return &PromosServiceImpl{PromosRepository: promosRepository, Validate: validate}
}

func (p PromosServiceImpl) FindAll(accountID int) []response.PromosResponse {
	result := p.PromosRepository.FindAll(accountID)

	var promos []response.PromosResponse
	for _, value := range result {
		promo := response.PromosResponse{
			PromoName:       value.PromoName,
			Image:           value.Image,
			Type:            value.Type,
			OutdatePromo:    value.OutdatePromo,
			Description:     value.Description,
			DetailTutorial:  value.DetailTutorial,
			DetailCondition: value.DetailCondition,
		}
		promos = append(promos, promo)
	}

	return promos
}

func (p PromosServiceImpl) FindRegular() []response.PromosResponse {
	result := p.PromosRepository.FindRegular()

	var promos []response.PromosResponse
	for _, value := range result {
		promo := response.PromosResponse{
			PromoName:       value.PromoName,
			Image:           value.Image,
			Type:            value.Type,
			OutdatePromo:    value.OutdatePromo,
			Description:     value.Description,
			DetailTutorial:  value.DetailTutorial,
			DetailCondition: value.DetailCondition,
		}
		promos = append(promos, promo)
	}

	return promos
}
