package service

import (
	"github.com/go-playground/validator/v10"
	"roastkuy-api/data/response"
	"roastkuy-api/repository"
)

type OutletsServiceImpl struct {
	OutletsRepository repository.OutletsRepository
	Validate          *validator.Validate
}

func NewOutletsServiceImpl(outletsRepository repository.OutletsRepository, validate *validator.Validate) OutletsService {
	return &OutletsServiceImpl{OutletsRepository: outletsRepository, Validate: validate}
}

func (o OutletsServiceImpl) FindAll() []response.OutletsResponse {
	result := o.OutletsRepository.FindAll()

	var outlets []response.OutletsResponse
	for _, value := range result {
		outlet := response.OutletsResponse{
			Id:                value.Id,
			OutletName:        value.OutletName,
			Slug:              value.Slug,
			Address:           value.Address,
			Latitude:          value.Latitude,
			Longitude:         value.Longitude,
			OperationTime:     value.OperationTime,
			Contact:           value.Contact,
			GofoodLink:        value.GofoodLink,
			ShopeefoodLink:    value.ShopeefoodLink,
			GrabfoodLink:      value.GrabfoodLink,
			TravelokaEatsLink: value.TravelokaEatsLink,
		}
		outlets = append(outlets, outlet)
	}

	return outlets
}

func (o OutletsServiceImpl) FindBySlug(slug string) (response.OutletsResponse, error) {
	outletData, err := o.OutletsRepository.FindBySlug(slug)
	if err != nil {
		return response.OutletsResponse{}, err
	}

	outletResponse := response.OutletsResponse{
		Id:                outletData.Id,
		OutletName:        outletData.OutletName,
		Slug:              outletData.Slug,
		Address:           outletData.Address,
		Latitude:          outletData.Latitude,
		Longitude:         outletData.Longitude,
		OperationTime:     outletData.OperationTime,
		Contact:           outletData.Contact,
		GofoodLink:        outletData.GofoodLink,
		ShopeefoodLink:    outletData.ShopeefoodLink,
		GrabfoodLink:      outletData.GrabfoodLink,
		TravelokaEatsLink: outletData.TravelokaEatsLink,
	}

	return outletResponse, nil
}
