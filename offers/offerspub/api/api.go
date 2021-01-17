package api

import (
	"net/http"

	"github.com/devminnu/assignments/offers/offerspub/controller"
	"github.com/devminnu/assignments/offers/offerspub/router"
	"github.com/devminnu/assignments/offers/offerspub/service"
	"github.com/gin-gonic/gin"
)

var (
	offersService    service.OffersService       = service.New()
	offersController controller.OffersController = controller.New(offersService)
)

//Init initialize api routes
func Init() {
	r := router.Get()
	r.GET("/publishOffers", func(ctx *gin.Context) {
		resp, err := offersController.PublishOffers(ctx)
		if err != nil {
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"status":   "FAILED",
				"error":    err,
				"response": resp,
			})
		}
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"status":   "SUCCESS",
			"error":    err,
			"response": resp,
		})
		// ctx.JSON(200, offersController.PublishOffers(ctx))
	})
}
