package database

import (
	"fmt"
	"github.com/Dencyuman/logvista-server/config"
	"github.com/Dencyuman/logvista-server/src/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.AppConfig.Host,
		config.AppConfig.User,
		config.AppConfig.Password,
		config.AppConfig.Dbname,
		config.AppConfig.Port,
		config.AppConfig.Sslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&models.Log{}, &models.Traceback{}, &models.System{})
	if err != nil {
		log.Println("Migration failed:", err)
		return err
	}
	return nil
}

func ResetTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(&models.Log{}, &models.Traceback{}, &models.System{})
	if err != nil {
		log.Println("Failed to reset tables:", err)
		return err
	}
	return nil
}
