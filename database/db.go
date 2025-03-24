package database

import (
	"log"

	"github.com/sohibjon7731/nectar/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DBConnect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetDBDSN()), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database" + err.Error())
		return nil, err
	}
	return db, nil
}
