package config

import (
	logger "gorestapi/logger"
	"github.com/spf13/viper"
)

var configuration Configuration
var	RouterPort int
var DatabaseType string
var DatabaseConnectionUri string
var CreateBodyLimit int64
var SecurityTokenSecret string

func init() {
	viper.SetConfigName("settings")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		logger.Log.Printf("Fatal error reading config file: %s \n", err)
		panic("config.init: Error reading config file")
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		logger.Log.Printf("Unable to decode into struct: %s \n", err)
		panic("config.init: Unable to decode into struct")
	}

	RouterPort = configuration.Router.Port
	DatabaseType = configuration.Database.Type
	DatabaseConnectionUri = configuration.Database.ConnectionUri
	CreateBodyLimit = configuration.Create.BodyLimit
	SecurityTokenSecret = configuration.Security.TokenSecret
}
