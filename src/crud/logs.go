package crud

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/Dencyuman/logvista-server/src/models"
	"github.com/Dencyuman/logvista-server/src/schemas"

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

// FindLogsとCountLogsのオプション型
type QueryLogsOptions struct {
	Limit       *int
	Offset      *int
	StartDate   *time.Time // 検索する日付の開始範囲
	EndDate     *time.Time // 検索する日付の終了範囲
	LevelName   *string    // ログレベル
	SystemId    *string    // システムID
	ContainsMsg *string    // メッセージ内容を含むかどうか
	ExcType     *string    // エラーの種類
	ExcDetail   *string    // エラーの詳細
	FileName    *string    // ファイル名
	Lineno      *int       // エラーが発生した行番号
}

// Logデータをデータベースから取得する
func FindLogs(db *gorm.DB, opts *QueryLogsOptions) ([]models.Log, error) {
	if opts == nil {
		opts = &QueryLogsOptions{}
	}

	query := db.Table("logs").Preload("System").Preload("ExcTraceback")

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

	if opts.SystemId != nil {
		query = query.Joins("JOIN systems ON systems.id = logs.system_id").
			Where("systems.id = ?", *opts.SystemId)
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

// Logデータをデータベースから取得する
func CountLogs(db *gorm.DB, opts *QueryLogsOptions) (int, error) {
	if opts == nil {
		opts = &QueryLogsOptions{}
	}

	query := db.Table("logs")

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

	if opts.SystemId != nil {
		query = query.Joins("JOIN systems ON systems.id = logs.system_id").
			Where("systems.id = ?", *opts.SystemId)
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

	var CountBuffer int64
	if err := query.Count(&CountBuffer).Error; err != nil {
		return 0, err
	}
	var Count int = int(CountBuffer)
	return Count, nil
}

// 指定したシステムの最新のログを取得する
func FindLatestLog(db *gorm.DB, systemName string) (*models.Log, error) {
	var log models.Log
	if err := db.Table("logs").Preload("ExcTraceback").Joins("JOIN systems ON systems.id = logs.system_id").
		Where("systems.name = ?", systemName).Order("timestamp desc").First(&log).Error; err != nil {
		return nil, err
	}
	return &log, nil
}

func FindSummaryData(
	db *gorm.DB,
	system *models.System,
	timeSpan int,
	latestTime time.Time,
	dataCount int,
) ([]schemas.SummaryData, error) {
	summaryData := make([]schemas.SummaryData, 0, dataCount)

	for i := 0; i < dataCount; i++ {
		// latestTimeからtimeSpan * dataCount秒前の時間を計算して、そこから集計を開始します。
		baseTime := latestTime.Add(time.Duration(-i*timeSpan) * time.Second)
		endTime := baseTime.Add(time.Duration(timeSpan) * time.Second)

		// 各ログレベルのカウントを取得
		var infoCount, warningCount, errorCount int64

		// INFOログカウント
		err := db.Model(&models.Log{}).
			Where("system_id = ? AND level_name = ? AND timestamp >= ? AND timestamp < ?", system.ID, "INFO", baseTime, endTime).
			Count(&infoCount).Error
		if err != nil {
			return nil, err
		}

		// WARNINGログカウント
		err = db.Model(&models.Log{}).
			Where("system_id = ? AND level_name = ? AND timestamp >= ? AND timestamp < ?", system.ID, "WARNING", baseTime, endTime).
			Count(&warningCount).Error
		if err != nil {
			return nil, err
		}

		// ERRORログカウント
		err = db.Model(&models.Log{}).
			Where("system_id = ? AND level_name = ? AND timestamp >= ? AND timestamp < ?", system.ID, "ERROR", baseTime, endTime).
			Count(&errorCount).Error
		if err != nil {
			return nil, err
		}

		// 計算した時間範囲のデータを追加
		summaryData = append(summaryData, schemas.SummaryData{
			BaseTime:        baseTime,
			InfologCount:    infoCount,
			WarninglogCount: warningCount,
			ErrorlogCount:   errorCount,
		})
	}

	return summaryData, nil
}
