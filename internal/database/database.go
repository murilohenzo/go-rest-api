package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Credentials struct {
	DbUsername string
	DbPassword string
	DbHost string
	DbTable string
	DbPort string
}

// NewDatabase - returns a pointer to a new database connection
func NewDatabase() (*gorm.DB, error) {
	fmt.Println("Setting up new database connection")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	c := Credentials{
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PORT"),
	}

	// dbConnectionString := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + strconv.Itoa(dbPort) + ")/" + dbTable
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.DbHost, c.DbPort, c.DbUsername, c.DbTable, c.DbPassword)
	fmt.Println(connectionString)

	db, err := gorm.Open("postgres", connectionString)
	if err != nil {
		return db, nil
	}

	if err := db.DB().Ping(); err != nil {
		return db, err
	}

	return db, nil
}