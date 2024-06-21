package config

import (
	"os"

	"github.com/jeffersonborges/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		// logger.Info("database file not found, creating...")
		// Create the database file and directory
		err := os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			// logger.Error(err)
			return nil, err
		}
		// Create the database file
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Error(err)
			return nil, err
		}
		// logger.Info("database file created successfully")
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		// logger.Errorf("sqlite opening erro: ", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		// logger.Errorf("sqlite automigration error: ", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}
