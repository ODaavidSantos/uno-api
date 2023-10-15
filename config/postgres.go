package config

import (
	"os"

	"github.com/DaviidSantos/uno-api/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Connect to database
	db, err := gorm.Open(postgres.Open(os.Getenv("DB")), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres initialization error: %v", err)
		return nil, err
	}

	// Migrate the schema
	err = db.AutoMigrate(&schemas.Solicitante{}, &schemas.Usuario{})
	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}

	return db, nil
}