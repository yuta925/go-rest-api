package db

import (
	"log"
	"os"

	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load() 
		if err != nil {
			log.Fatalln(err)
		}
	}

	url := fmt.Sprintf("postgress://%s:%s:%s:%s:%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"),os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	
	if err != nil {
		log.Fatalln(err)
	} 
	fmt.Println("Conected")
	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, _ := db.DB()
	if err := sqlDB.Close(); err != nil {
		log.Fatalln(err)
	}
}