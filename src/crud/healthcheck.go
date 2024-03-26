package crud

import (
	"errors"
	"github.com/Dencyuman/logvista-server/src/models"
	"gorm.io/gorm"
	"log"
	"strings"
)

func InsertHealthcheck(db *gorm.DB, modelHealthcheckConfig *models.HealthcheckConfig) error {
	if modelHealthcheckConfig == nil {
		return errors.New("received nil healthcheck config data")
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
	// Logデータをデータベースに挿入
	if err := tx.Create(modelHealthcheckConfig).Error; err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			// 一意制約違反のエラーの場合、ログを出力してロールバック
			log.Printf("Warning: duplicate key value violates unique constraint: %v\n", err)
			tx.Rollback()
			return nil // ここでnilを返すことで、エラーとして扱わない
		} else {
			// その他のエラーの場合、ロールバックしてエラーを返す
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error // トランザクションをコミット
}

func FindHealthcheckConfigs(db *gorm.DB, systemID string) ([]models.HealthcheckConfig, error) {
	var healthcheckConfigs []models.HealthcheckConfig
	if err := db.Where("system_id = ?", systemID).Find(&healthcheckConfigs).Error; err != nil {
		return nil, err
	}
	return healthcheckConfigs, nil
}

func FindAllActiveHealthcheckConfigs(db *gorm.DB) ([]models.HealthcheckConfig, error) {
	var healthcheckConfigs []models.HealthcheckConfig
	if err := db.Where("is_active = ?", true).Find(&healthcheckConfigs).Error; err != nil {
		return nil, err
	}
	return healthcheckConfigs, nil
}

// InsertHealthcheckLog は新しいHealthcheckLogエントリをデータベースに挿入します。
func InsertHealthcheckLog(db *gorm.DB, hl *models.HealthcheckLog) error {
	err := db.Create(hl).Error
	if err != nil {
		return err
	}
	return nil
}

func FindHealthcheckLogs(db *gorm.DB, systemID string) ([]models.HealthcheckLog, error) {
	var healthcheckLogs []models.HealthcheckLog
	if err := db.Where("system_id = ?", systemID).Find(&healthcheckLogs).Error; err != nil {
		return nil, err
	}
	return healthcheckLogs, nil
}
