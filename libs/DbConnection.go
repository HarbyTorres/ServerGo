package libs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (config *DBconfig) InitConnection() *gorm.DB {
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s  sslmode=disable", config.user, config.password, config.host, config.port, config.database)
	db, err := gorm.Open(postgres.Open(connString))
	if err != nil {
		log.Panic(err)
		os.Exit(-1)
	}

	return db
}
