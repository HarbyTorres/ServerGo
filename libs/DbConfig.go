package libs

import (
	"log"

	config "github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBconfig struct {
	host     string
	port     string
	database string
	user     string
	password string
}

func Configure(configPath string, configName string) DBconfig {
	config.AddConfigPath(configPath)
	config.SetConfigName(configName)
	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file %s", err)
	}
	response := DBconfig{
		config.GetString("default.host"),
		config.GetString("default.port"),
		config.GetString("default.database"),
		config.GetString("default.user"),
		config.GetString("default.password"),
	}

	return response
}
