package consumer

import (
	"encoding/json"
	"fmt"

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

	/* 	db.DropTable(&User{})
	   	db.DropTable(&Bike{})
	   	db.DropTable(&Company{})
	   	//
	   	db.CreateTable(&User{})
	   	db.CreateTable(&Bike{})
	   	db.CreateTable(&Company{})

	   	db.Debug().Create(&Bike{
	   		BikeID: "c321",
	   		Number: "321234",
	   		User: User{
	   			UserID: "121",
	   			Name:   "Minhaj",
	   			Age:    27,
	   		},
	   		Company: Company{
	   			CompanyID: "1578",
	   			Name:      "MKCL",
	   			Location:  "Pune",
	   		},
	   	})
	*/
	// db.CreateTable(&model.User{})
	// db.CreateTable(&model.CreditCard{})
	// db.Debug().Create(&model.User{
	// 	UserID: "111",
	// 	Name:   "Minhaj",
	// 	Age:    27,
	// 	CreditCard: []model.CreditCard{
	// 		{
	// 			Number: "123",
	// 		},
	// 		{
	// 			Number: "456",
	// 		},
	// 	},
	// })
	for _, offer := range offers.Offers {
		// fmt.Println(offer.Hotel)
		// j, _ := json.Marshal(offer)
		// fmt.Println(string(j))
		// h := offer.Hotel
		// h.Offers = []*model.Offer{&offer}
		// offer.Room.Offers = []*model.Offer{&offer}
		// h.Room = []model.Room{offer.Room}
		// h.RatePlan = offer.RatePlan
		// db.Debug().Save(&h)
		fmt.Println("lo", offer.Fees)
		// fmt.Printf("%#v", offer.Fees)
		db.Debug().Create(&offer)
	}
}

/* type Bike struct {
	// gorm.Model
	BikeID    string //`gorm:"primary_key"`
	Number    string
	Type      string
	UserID    string
	User      User `gorm:"foreignKey:UserID;references:UserID"`
	CompanyID string
	Company   Company `gorm:"foreignKey:CompanyID;references:CompanyID"`
}

type User struct {
	// gorm.Model
	UserID string `gorm:"primary_key"`
	Name   string
	Age    int
}

type Company struct {
	CompanyID string `gorm:"primary_key"`
	Name      string
	Location  string
} */
