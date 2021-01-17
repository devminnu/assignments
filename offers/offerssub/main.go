package main

import (
	"github.com/devminnu/assignments/offers/offerssub/config"
	"github.com/devminnu/assignments/offers/offerssub/consumer"
	"github.com/devminnu/assignments/offers/offerssub/db"
	"github.com/devminnu/assignments/offers/offerssub/logger"
)

func main() {
	config.Load() //Load all config variable and values
	logger.Init() //Initialize logger
	db.Init()     //MY_SQL database initialization
	// defer db.Get().Close() //Close db connection
	consumer.Init()
	defer consumer.GetChannel().Close() //Close RABBIT_MQ connection channel
	defer consumer.GetConnection().Close()
	consumer.SubscribeOffers()
}
