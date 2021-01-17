package controller

import (
	log "github.com/devminnu/assignments/offers/offerspub/logger"
	"github.com/devminnu/assignments/offers/offerspub/model"
	"github.com/devminnu/assignments/offers/offerspub/service"
	"github.com/gin-gonic/gin"
)

type OffersController interface {
	PublishOffers(ctx *gin.Context) (model.Offers, error)
}

type controller struct {
	service service.OffersService
}

func New(service service.OffersService) OffersController {
	return &controller{
		service: service,
	}
}

func (c *controller) PublishOffers(ctx *gin.Context) (model.Offers, error) {
	var offers model.Offers
	err := ctx.BindJSON(&offers)
	if err != nil {
		log.Get().Error(err)
		return offers, err
	}
	return c.service.PublishOffers(offers)
}
