package db

import (
	"errors"

	"github.com/devminnu/assignments/offers/offerssub/config"
	log "github.com/devminnu/assignments/offers/offerssub/logger"
	"github.com/devminnu/assignments/offers/offerssub/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Init initializes db confection
func Init() {
	var err error
	dbhost := config.ReadEnvString("MYSQL_ADDON_HOST")
	dbname := config.ReadEnvString("MYSQL_ADDON_DB")
	dbuser := config.ReadEnvString("MYSQL_ADDON_USER")
	dbpassword := config.ReadEnvString("MYSQL_ADDON_PASSWORD")
	dbport := config.ReadEnvString("MYSQL_ADDON_PORT")

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Get().Error("DB_CONNECTION_FAILED", err)
		panic(err)
	} else {
		log.Get().Info("DB_CONNECTION_SUCCESSFULL")
	}
	db.DropTable(&model.Offer{})
	db.DropTable(&model.Hotel{})
	db.DropTable(&model.Room{})
	db.DropTable(&model.RatePlan{})
	db.DropTable(&model.Amenity{})
	db.DropTable(&model.CancellationPolicy{})
	db.DropTable(&model.OtherCondition{})
	db.DropTable(&model.Fee{})
	//
	db.CreateTable(&model.Offer{})
	db.CreateTable(&model.Hotel{})
	db.CreateTable(&model.Room{})
	db.CreateTable(&model.RatePlan{})
	db.CreateTable(&model.Amenity{})
	db.CreateTable(&model.CancellationPolicy{})
	db.CreateTable(&model.OtherCondition{})
	db.CreateTable(&model.Fee{})
}

func Get() *gorm.DB {
	if db == nil {
		log.Get().Error("DB_NOT_INITIALIZED")
		panic(errors.New("DB_NOT_INITIALIZED"))
	}
	return db
}
