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

// Systemデータをデータベースに更新する
func UpdateSystem(db *gorm.DB, modelSystem *models.System) error {
	if modelSystem == nil {
		return errors.New("received nil system data")
	}

	// Begin a new transaction.
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// Update the modelSystem. Assuming you want to update all fields except primary key.
	if err := tx.Model(modelSystem).Updates(modelSystem).Error; err != nil {
		tx.Rollback()
		return err
	}

	// If no error, commit the transaction.
	return tx.Commit().Error
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

// SystemデータをIDで検索する
func FindSystemByID(db *gorm.DB, id string) (*models.System, error) {
	var system models.System
	if err := db.Where("id = ?", id).First(&system).Error; err != nil {
		return nil, err
	}
	return &system, nil
}

// Systemデータをデータベースから削除する
func DeleteSystem(db *gorm.DB, systemId string) error {
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// システムに紐づく各LogデータのTracebackデータを削除する
	if err := tx.Where("log_id IN (SELECT id FROM logs WHERE system_id = ?)", systemId).Delete(&models.Traceback{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// システムに紐づくLogデータを削除する
	if err := tx.Where("system_id = ?", systemId).Delete(&models.Log{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// システムデータを削除する
	if err := tx.Where("id = ?", systemId).Delete(&models.System{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
