package consumer

import (
	"github.com/devminnu/assignments/offers/offerssub/config"
	log "github.com/devminnu/assignments/offers/offerssub/logger"
	"github.com/streadway/amqp"
)

var (
	connection *amqp.Connection
	channel    *amqp.Channel
)

func Init() {
	var err error
	rbmqHost := config.ReadEnvString("RABBIT_MQ_HOST")
	rbmqPort := config.ReadEnvString("RABBIT_MQ_PORT")
	rbmqUser := config.ReadEnvString("RABBIT_MQ_USER")
	rbmqPassword := config.ReadEnvString("RABBIT_MQ_PASSWORD")

	rbmqDSN := "amqp://" + rbmqUser + ":" + rbmqPassword + "@" + rbmqHost + ":" + rbmqPort + "/"
	connection, err = amqp.Dial(rbmqDSN)
	if err != nil {
		log.Get().Error(err)
		return
	}

	channel, err = connection.Channel()
	if err != nil {
		log.Get().Error(err)
		return
	}
}

func InitSubscribeOffers() {}

func GetChannel() *amqp.Channel {
	if channel == nil {
		log.Get().Error("RABBIT_MQ_CHANNEL_NOT_INITIALIZED")
		return nil
	}
	return channel
}
func GetConnection() *amqp.Connection {
	if connection == nil {
		log.Get().Error("RABBIT_MQ_CONNECTION_NOT_INITIALIZED")
		return nil
	}
	return connection
}
