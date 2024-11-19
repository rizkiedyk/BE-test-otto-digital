package main

import (
	"test-ottodigital-be/config"
	"test-ottodigital-be/router"

	"github.com/op/go-logging"
	"github.com/spf13/viper"
)

var logger = logging.MustGetLogger("main")

func init() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config.ConfigureLogger()

}

func main() {
	router := router.SetupRouter()

	router.Run(":1010")
}
