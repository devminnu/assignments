package pub

import (
	"encoding/json"
	"errors"

	"github.com/streadway/amqp"

	log "github.com/devminnu/assignments/offers/offerspub/logger"
	"github.com/devminnu/assignments/offers/offerspub/model"
)

func PublishOffers(offers model.Offers) (model.Offers, error) {
	if len(offers.Offers) == 0 {
		return model.Offers{}, errors.New("NO_OFFERS")
	}
	ch := GetChannel()
	if ch == nil {
		return model.Offers{}, errors.New("RABBIT_MQ_CHANNEL_NOT_INITIALIZED")
	}
	err := ch.ExchangeDeclare(
		"publishOffers", // name
		"fanout",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // no-wait
		nil,             // arguments
	)
	if err != nil {
		log.Get().Error(err)
		return offers, err
	}
	offersBS, err := json.Marshal(offers)
	if err != nil {
		log.Get().Error(err)
		return offers, err
	}
	err = ch.Publish(
		"publishOffers", // exchange
		"",              // routing key
		false,           // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(offersBS),
		})
	if err != nil {
		log.Get().Error(err)
		return offers, err
	}
	return offers, nil
}
