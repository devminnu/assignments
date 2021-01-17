package main

import (
	"fmt"
	"strconv"

	"github.com/devminnu/assignments/offers/offerspub/api"
	"github.com/devminnu/assignments/offers/offerspub/config"
	"github.com/devminnu/assignments/offers/offerspub/db"
	"github.com/devminnu/assignments/offers/offerspub/logger"
	"github.com/devminnu/assignments/offers/offerspub/pub"
	"github.com/devminnu/assignments/offers/offerspub/router"
)

func main() {
	config.Load()          //Load all config variable and values
	logger.Init()          //Initialize logger
	db.Init()              //MY_SQL database initialization
	defer db.Get().Close() //Close db connection
	pub.Init()
	defer pub.GetChannel().Close() //Close RABBIT_MQ connection channel
	defer pub.GetConnection().Close()
	router.Init()
	api.Init()
	router := router.Get()
	port := config.AppPort() // This should be changed to the service port number via argument or environment variable.
	addr := fmt.Sprintf(":%s", strconv.Itoa(port))
	router.Run(addr)
}
