package crud

import (
	"errors"
	"log"
	"strings"
	"time"

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
	Limit       *int
	Offset      *int
	StartDate   *time.Time // 検索する日付の開始範囲
	EndDate     *time.Time // 検索する日付の終了範囲
	LevelName   *string    // ログレベル
	SystemName  *string    // システム名
	ContainsMsg *string    // メッセージ内容を含むかどうか
	ExcType     *string    // エラーの種類
	ExcDetail   *string	   // エラーの詳細
	FileName    *string    // ファイル名
	Lineno      *int       // エラーが発生した行番号
}

// Logデータをデータベースから取得する
func FindLogs(db *gorm.DB, opts *FindLogsOptions) ([]models.Log, error) {
	if opts == nil {
		opts = &FindLogsOptions{}
	}

	query := db.Table("logs").Preload("ExcTraceback")

	if opts.Limit != nil {
		query = query.Limit(*opts.Limit)
	}

	if opts.Offset != nil {
		query = query.Offset(*opts.Offset)
	}

	if opts.StartDate != nil {
		query = query.Where("timestamp >= ?", *opts.StartDate)
	}

	if opts.EndDate != nil {
		query = query.Where("timestamp <= ?", *opts.EndDate)
	}

	if opts.LevelName != nil {
		query = query.Where("level_name = ?", *opts.LevelName)
	}

	if opts.SystemName != nil {
		query = query.Where("system_name = ?", *opts.SystemName)
	}

	if opts.ContainsMsg != nil {
		query = query.Where("message LIKE ?", "%"+*opts.ContainsMsg+"%")
	}

	if opts.ExcType != nil {
		query = query.Where("exc_type = ?", *opts.ExcType)
	}

	if opts.ExcDetail != nil {
		query = query.Where("exc_detail LIKE ?", "%"+*opts.ExcDetail+"%")
	}

	if opts.FileName != nil {
		query = query.Where("file_name = ?", *opts.FileName)
	}

	if opts.Lineno != nil {
		query = query.Where("lineno = ?", *opts.Lineno)
	}

	var logs []models.Log
	if err := query.Find(&logs).Error; err != nil {
		return nil, err
	}
	return logs, nil
}
