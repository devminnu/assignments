package service

import (
	"github.com/devminnu/assignments/offers/offerspub/model"
	"github.com/devminnu/assignments/offers/offerspub/pub"
)

type OffersService interface {
	PublishOffers(model.Offers) (model.Offers, error)
}

// type offersService model.Offers
type offersService struct {
	Offers model.Offers
}

func New() OffersService {
	return &offersService{
		Offers: model.Offers{},
	}
}

func (service *offersService) PublishOffers(offers model.Offers) (model.Offers, error) {
	return pub.PublishOffers(offers)
}
