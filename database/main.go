package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	logger "gorestapi/logger"
	config "gorestapi/config"
)

type DBManager struct {
	db *gorm.DB
}

var dbManagerInstance = new()

func new() DBManager {
	logger.Log.Println("Connecting to database")
	localDbRef, err := gorm.Open(config.DatabaseType, config.DatabaseConnectionUri)
	if err != nil {
		logger.Log.Printf("Failed to connect to database: %s \n", err)
		panic("database.openDB: Failed to connect to database")
	} else {
		logger.Log.Println("Database Initialized successfully")
	}
	return DBManager{db: localDbRef}
}

func CloseDB() {
	dbManagerInstance.db.Close()
	logger.Log.Println("Database closed gracefully")
}

func GetDB() *gorm.DB {
	return dbManagerInstance.db
}
