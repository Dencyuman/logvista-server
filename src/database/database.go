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
		config.AppConfig.DbHost,
		config.AppConfig.DbUser,
		config.AppConfig.DbPassword,
		config.AppConfig.Dbname,
		config.AppConfig.DbPort,
		config.AppConfig.DbSslmode,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Log{},
		&models.Traceback{},
		&models.System{},
		&models.HealthcheckConfig{},
		&models.HealthcheckLog{},
	)
	if err != nil {
		log.Println("Migration failed:", err)
		return err
	}
	return nil
}

func ResetTables(db *gorm.DB) error {
	err := db.Migrator().DropTable(
		&models.Log{},
		&models.Traceback{},
		&models.System{},
		&models.HealthcheckConfig{},
		&models.HealthcheckLog{},
	)
	if err != nil {
		log.Println("Failed to reset tables:", err)
		return err
	}
	return nil
}
