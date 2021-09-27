package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/vitorbgouveia/go-web-api/configs/database/migrations"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type connectionDatabase struct {
	host     string
	name     string
	user     string
	password string
	port     int
	timeZone string
}

func getConnection() connectionDatabase {
	port, _ := strconv.Atoi(os.Getenv("DATABASE_PORT"))

	return connectionDatabase{
		host:     os.Getenv("DATABASE_HOST"),
		name:     os.Getenv("DATABASE_NAME"),
		user:     os.Getenv("DATABASE_USER"),
		password: os.Getenv("DATABASE_PASSWORD"),
		port:     port,
		timeZone: os.Getenv("DATABASE_TIME_ZONE"),
	}
}

// StartDB up and configure Database
func StartDB() {
	var connection connectionDatabase = getConnection()
	str := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s",
		connection.host, connection.user, connection.password, connection.name, connection.port, connection.timeZone)

	database, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("Database Error: ", err)
	}

	db = database

	config, _ := db.DB()

	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)
	migrations.RunMigrations(db)
}

// GetDatabase return connection database
func GetDatabase() *gorm.DB {
	return db
}
