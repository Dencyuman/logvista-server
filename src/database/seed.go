package database

import (
	"encoding/json"
	"fmt"
	"github.com/Dencyuman/logvista-server/src/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

// Systemテーブルにシードする
func seedSystemTable(db *gorm.DB) error {
	systems := []models.System{
		{
			Name:     "systemA",
			Category: "API Server",
		},
		{
			Name:     "systemB",
			Category: "API Server",
		},
		{
			Name:     "systemC",
			Category: "Automation",
		},
		{
			Name:     "systemD",
			Category: "Automation",
		},
		{
			Name:     "systemE",
			Category: "Web Server",
		},
		{
			Name:     "systemF",
			Category: "Web Server",
		},
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, system := range systems {
		if err := tx.Create(&system).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// generateRandomLogはランダムなログデータを生成する
func generateRandomLog() *models.Log {
	// SystemNameの選択肢
	systemNames := []string{"systemA", "systemB", "systemC", "systemD", "systemE", "systemF"}
	// LevelNameの選択肢
	levelNames := []string{"INFO", "WARNING", "ERROR"}

	// 現在時刻から過去24時間のタイムスタンプを生成
	randomTimestamp := time.Now().Add(-time.Duration(rand.Intn(6)) * time.Hour)

	// Attributesをマップとして生成
	randomAttributes := make(map[string]interface{})
	for i := 0; i < rand.Intn(5)+1; i++ { // 1から5までのランダムな数の属性を生成
		key := fmt.Sprintf("key%d", i) // 例: "key1", "key2", ...
		value := rand.Intn(100)        // 値は0から99までのランダムな整数
		randomAttributes[key] = value
	}
	jsonAttributes, err := json.Marshal(randomAttributes)
	if err != nil {
		// エラーハンドリング
		fmt.Println("Error marshaling attributes:", err)
		return nil
	}

	// Logオブジェクトを生成
	log := &models.Log{
		ID:              uuid.New().String(),
		SystemName:      systemNames[rand.Intn(len(systemNames))],
		CPUPercent:      rand.Float64() * 100,
		ExcType:         "ExceptionType",
		ExcValue:        "ExceptionValue",
		ExcDetail:       "ExceptionDetail",
		FileName:        "file.py",
		FuncName:        "funcName",
		Lineno:          rand.Intn(1000),
		Message:         "Random log message",
		Module:          "main",
		Name:            "RandomLogger",
		LevelName:       levelNames[rand.Intn(len(levelNames))],
		Levelno:         rand.Intn(100),
		Process:         rand.Intn(10000),
		ProcessName:     "RandomProcess",
		Thread:          rand.Intn(10000),
		ThreadName:      "RandomThread",
		TotalMemory:     int64(rand.Intn(10000000)),
		AvailableMemory: int64(rand.Intn(10000000)),
		MemoryPercent:   rand.Float64() * 100,
		UsedMemory:      int64(rand.Intn(10000000)),
		FreeMemory:      int64(rand.Intn(10000000)),
		CPUUserTime:     rand.Float64() * 100,
		CPUSystemTime:   rand.Float64() * 100,
		CPUIdleTime:     rand.Float64() * 100,
		Timestamp:       randomTimestamp,
		Attributes:      string(jsonAttributes),
		CreatedAt:       randomTimestamp,
		UpdatedAt:       randomTimestamp,
	}

	// LevelNameがERRORの場合にのみTracebackを追加
	if log.LevelName == "ERROR" {
		// Tracebackを1つ以上追加したい場合は、このループの回数を変更します
		for i := 0; i < rand.Intn(5)+1; i++ {
			traceback := generateRandomTraceback(log.ID, log.Timestamp)
			log.ExcTraceback = append(log.ExcTraceback, *traceback)
		}
	}

	return log
}

// generateRandomTracebackはランダムなTracebackデータを生成します。
func generateRandomTraceback(logID string, logTimestamp time.Time) *models.Traceback {
	traceback := &models.Traceback{
		LogID:      logID,
		TbFilename: "file.py",
		TbLineno:   rand.Intn(1000),
		TbName:     "TracebackFunctionName",
		TbLine:     "Some random code line",
		CreatedAt:  logTimestamp,
		UpdatedAt:  logTimestamp,
	}
	return traceback
}

// Logテーブルにシードする
func seedLogTable(db *gorm.DB) error {
	// ランダムなログを1000件生成
	logs := make([]*models.Log, 0)
	for i := 0; i < 1000; i++ {
		logs = append(logs, generateRandomLog())
	}

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	for _, log := range logs {
		if err := tx.Create(log).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit().Error
}

// 全てのSeedingを実行する
func Seed(db *gorm.DB) error {
	if err := ResetTables(db); err != nil {
		return err
	}
	if err := Migrate(db); err != nil {
		return err
	}
	if err := seedSystemTable(db); err != nil {
		return err
	}
	if err := seedLogTable(db); err != nil {
		return err
	}
	return nil
}
