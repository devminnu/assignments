package consumer

import (
	"encoding/json"

	log "github.com/devminnu/assignments/offers/offerssub/logger"

	"github.com/devminnu/assignments/offers/offerssub/db"
	"github.com/devminnu/assignments/offers/offerssub/model"
	"github.com/sirupsen/logrus"
)

func SubscribeOffers() {
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
		return
	}
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Get().Error(err)
		return
	}
	err = ch.QueueBind(
		q.Name,          // queue name
		"",              // routing key
		"publishOffers", // exchange
		false,
		nil)
	if err != nil {
		log.Get().Error(err)
		return
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Get().Error(err)
		return
	}

	forever := make(chan bool)

	var offers model.Offers
	go func() {
		for d := range msgs {
			err := json.Unmarshal(d.Body, &offers)
			if err != nil {
				logrus.Error(err)
			}
			ProcessOffers(offers)
			d.Ack(false)
		}
	}()

	log.Get().Print("[*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func ProcessOffers(offers model.Offers) {
	db := db.Get()
	for _, offer := range offers.Offers {
		tx := db.Begin()
		if err := tx.Create(&offer).Error; err != nil {
			tx.Rollback()
		}
		tx.Commit()
	}
}
