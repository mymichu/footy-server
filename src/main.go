package main

import (
	"os"
	"footy-server-app/restapi"
	"footy-server-app/config"
	"log"
	"github.com/spf13/viper"
)

func main() {
	configFilefolder := os.Getenv("CONFIG_FILE_FOLDER")

	if len(configFilefolder) == 0 {
		log.Fatalf("CONFIG_FILE_FOLDER Environment not set!")
	}
	viper.SetConfigName("config")
	viper.AddConfigPath(configFilefolder)
	var configuration config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	log.Printf("database uri is %s", configuration.Database.ConnectionUri)
	log.Printf("port for this application is %d", configuration.Server.Port)
	
	restapi.Listen(configuration.Server.Port)
}