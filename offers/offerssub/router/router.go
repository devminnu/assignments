package router

import (
	"errors"

	"github.com/devminnu/assignments/offers/offerssub/logger"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

var r *gin.Engine

func Init() {
	r = gin.New()
	r.Use(ginlogrus.Logger(logger.Get()), gin.Recovery())
}

func Get() *gin.Engine {
	if r == nil {
		logrus.Error("GIN_ROUTER_NOT_INITIALIZED")
		panic(errors.New("GIN_ROUTER_NOT_INITIALIZED"))
	}
	return r
}
