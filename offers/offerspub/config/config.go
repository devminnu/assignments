package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

var (
	appName string
	appPort int
)

func init() {
	if os.Getenv("GIN_MODE") != "release" {
		os.Setenv("APP_NAME", "offerspub")
		os.Setenv("APP_PORT", "30003")
		os.Setenv("APP_LOG_PATH", "logs/log.log")
		//MYSQL DB
		os.Setenv("MYSQL_ADDON_HOST", "localhost")
		os.Setenv("MYSQL_ADDON_DB", "offers")
		os.Setenv("MYSQL_ADDON_USER", "admin")
		os.Setenv("MYSQL_ADDON_PASSWORD", "admin")
		os.Setenv("MYSQL_ADDON_PORT", "3306")
		//RabbitMQ
		os.Setenv("RABBIT_MQ_HOST", "localhost")
		os.Setenv("RABBIT_MQ_PORT", "5672")
		os.Setenv("RABBIT_MQ_USER", "guest")
		os.Setenv("RABBIT_MQ_PASSWORD", "guest")
	}
}

func Load() {
	viper.AutomaticEnv()
}

func AppName() string {
	if appName == "" {
		appName = ReadEnvString("APP_NAME")
	}
	fmt.Println("Hello")
	return appName
}

func AppPort() int {
	if appPort == 0 {
		appPort = ReadEnvInt("APP_PORT")
	}
	return appPort
}

func ReadEnvInt(key string) int {
	checkIfSet(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid integer", key))
	}
	return v
}

func ReadEnvString(key string) string {
	checkIfSet(key)
	return viper.GetString(key)
}

func ReadEnvBool(key string) bool {
	checkIfSet(key)
	return viper.GetBool(key)
}

func checkIfSet(key string) {
	if !viper.IsSet(key) {
		err := errors.New(fmt.Sprintf("Key %s is not set", key))
		panic(err)
	}
}
