package main

import (
	"github.com/mymichu/footy-server/internal/mock"
	"os"
	"github.com/mymichu/footy-server/internal/restapi"
	"github.com/mymichu/footy-server/internal/config"
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
	
	rest := restapi.RestAPISettings{mock.MockLogic{}}
	rest.Listen(configuration.Server.Port)
}