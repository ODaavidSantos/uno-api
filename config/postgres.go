package config

import (
	"os"

	"github.com/DaviidSantos/uno-api/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func initializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Load enviroment variables
	err := godotenv.Load(".env")
	if err != nil {
		logger.Errorf("dotenv load error: %v", err)
		return nil, err
	}

	// Connect to database
	db, err := gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres initialization error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Solicitante{})
	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}

	return db, nil
}