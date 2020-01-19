package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func NewDatabase() *gorm.DB {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config.env file: %v", err)
	}

	db, err := initDatabase(config)
	if err != nil {
		log.Fatalf("Error initializing the database: %v", err)
	}

	//if err := db.Ping(); err != nil {
	//	log.Fatalf("Error pinging DB: %v", err)
	//}

	return db
}

// Config represents structure of the config.env
type Config struct {
	dbUser string
	//dbPass string
	dbName string
	dbHost string
	dbPort string
}

func loadConfig() (config *Config, err error) {
	err = godotenv.Load("../config.env")
	if err != nil {
		log.Fatal("Error loading config.env file")
	}

	config = &Config{
		dbUser: os.Getenv("db_user"),
		//dbPass : os.Getenv("db_pass"),
		dbName: os.Getenv("db_name"),
		dbHost: os.Getenv("db_host"),
		dbPort: os.Getenv("db_port"),
	}
	return config, err
}

func initDatabase(c *Config) (db *gorm.DB, err error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		" dbname=%s sslmode=disable",
		c.dbHost, c.dbPort, c.dbUser, c.dbName)

	db, err = gorm.Open("postgres", psqlInfo)
	return db, err
}
