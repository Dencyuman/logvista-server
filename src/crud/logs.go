package crud

import (
	"errors"
	"log"
	"strings"

	"github.com/Dencyuman/logvista-server/src/models"
	"gorm.io/gorm"
)

// Logデータとそれに関連するTracebackデータをデータベースに挿入する
func InsertLogWithTracebacks(db *gorm.DB, modelLog *models.Log) error {
	if modelLog == nil {
		return errors.New("received nil log data")
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
	if err := tx.Create(modelLog).Error; err != nil {
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

// FindLogsのオプション型
type FindLogsOptions struct {
	Limit  *int
	Offset *int
}

// Logデータをデータベースから取得する
func FindLogs(db *gorm.DB, opts *FindLogsOptions) ([]models.Log, error) {
	if opts == nil {
		opts = &FindLogsOptions{}
	}

	// Preloadでリレーション先のトレースバックも取得
	query := db.Table("logs").Preload("ExcTraceback")

	if opts.Limit != nil {
		query = query.Limit(*opts.Limit)
	}

	if opts.Offset != nil {
		query = query.Offset(*opts.Offset)
	}

	var logs []models.Log
	if err := query.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
