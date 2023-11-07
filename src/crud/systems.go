package crud

import (
	"errors"

	"github.com/Dencyuman/logvista-server/src/models"
	"gorm.io/gorm"
)

// Systemデータをデータベースに挿入する
func InsertSystem(db *gorm.DB, modelSystem *models.System) error {
	if modelSystem == nil {
		return errors.New("received nil system data")
	}

	// トランザクションを開始
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Systemデータをデータベースに挿入
	if err := tx.Create(modelSystem).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error // トランザクションをコミット
}

// Systemデータを全件取得する
func FindAllSystems(db *gorm.DB) ([]models.System, error) {
	var systems []models.System
	if err := db.Find(&systems).Error; err != nil {
		return nil, err
	}
	return systems, nil
}

// Systemデータをシステム名で検索する
func FindSystemByName(db *gorm.DB, name string) (*models.System, error) {
	var system models.System
	if err := db.Where("name = ?", name).First(&system).Error; err != nil {
		return nil, err
	}
	return &system, nil
}