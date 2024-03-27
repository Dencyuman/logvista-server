package crud

import (
	"errors"
	"github.com/Dencyuman/logvista-server/src/models"
	"gorm.io/gorm"
	"log"
	"strings"
)

// InsertHealthcheck は新しいHealthcheckConfigエントリをデータベースに挿入します。
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

// FindHealthcheckConfigs は指定されたシステムに関連付けられたHealthcheckConfigエントリをデータベースから取得します。
func FindHealthcheckConfigs(db *gorm.DB, systemID string) ([]models.HealthcheckConfig, error) {
	var healthcheckConfigs []models.HealthcheckConfig
	if err := db.Where("system_id = ?", systemID).Find(&healthcheckConfigs).Error; err != nil {
		return nil, err
	}
	return healthcheckConfigs, nil
}

// FindAllActiveHealthcheckConfigs はデータベースからすべてのアクティブなHealthcheckConfigエントリを取得します。
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

// FindHealthcheckLogs は指定されたHealthcheckConfigに関連付けられたHealthcheckLogエントリをデータベースから取得します。
// count は取得するログの最大件数を指定します。desc が true の場合は結果を降順にソートします。
func FindHealthcheckLogs(db *gorm.DB, configID string, count int, desc bool) ([]models.HealthcheckLog, error) {
	var healthcheckLogs []models.HealthcheckLog
	query := db.Where("healthcheck_config_id = ?", configID)

	if desc {
		query = query.Order("created_at DESC")
	} else {
		query = query.Order("created_at ASC")
	}

	if count > 0 {
		query = query.Limit(count)
	}

	if err := query.Find(&healthcheckLogs).Error; err != nil {
		return nil, err
	}

	return healthcheckLogs, nil
}

// FindHealthcheckConfigByID は指定されたIDに一致するHealthcheckConfigエントリをデータベースから取得します。
func FindHealthcheckConfigByID(db *gorm.DB, id string) (*models.HealthcheckConfig, error) {
	var healthcheckConfig models.HealthcheckConfig
	if err := db.Where("id = ?", id).First(&healthcheckConfig).Error; err != nil {
		return nil, err
	}
	return &healthcheckConfig, nil
}
