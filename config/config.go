package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	DB *gorm.DB
)

func Connect() {
	d, err := gorm.Open("postgres", "host=localhost user=jhenriquepg dbname=usermanagement sslmode=disable password=1234")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = d
}

func GetDB() *gorm.DB {
	return DB
}
