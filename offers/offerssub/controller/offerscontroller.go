package controller

// import (
// 	"github.com/devminnu/assignments/offers/offerssub/model"
// 	"github.com/devminnu/assignments/offers/offerssub/service"
// 	"github.com/gin-gonic/gin"
// )

// type OffersController interface {
// 	PublishOffers(ctx *gin.Context) model.Offers
// }

// type controller struct {
// 	service service.OffersService
// }

// func New(service service.OffersService) OffersController {
// 	return &controller{
// 		service: service,
// 	}
// }

// func (c *controller) PublishOffers(ctx *gin.Context) model.Offers {
// 	var offers model.Offers
// 	ctx.BindJSON(&offers)
// 	c.service.PublishOffers(offers)
// 	return offers
// }
