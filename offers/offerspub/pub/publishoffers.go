package pub

import (
	"encoding/json"

	log "github.com/devminnu/assignments/offers/offerspub/logger"

	"github.com/streadway/amqp"

	"github.com/devminnu/assignments/offers/offerspub/model"
)

func PublishOffers(offers model.Offers) (model.Offers, error) {
	ch := GetChannel()
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
