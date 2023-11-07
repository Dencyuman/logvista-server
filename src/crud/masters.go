package crud

import (
	"gorm.io/gorm"
)

// DB上に存在するエラー型名一覧を取得する
func FindErrorTypes(db *gorm.DB) ([]string, error) {
	var errorTypes []string
	if err := db.Table("logs").Distinct("exc_type").Pluck("exc_type", &errorTypes).Error; err != nil {
		return nil, err
	}
	return errorTypes, nil
}

// DB上に存在するファイル名一覧を取得する
func FindFiles(db *gorm.DB) ([]string, error) {
	var files []string
	if err := db.Table("logs").Distinct("file_name").Pluck("file_name", &files).Error; err != nil {
		return nil, err
	}
	return files, nil
}

// DB上に存在するログレベル一覧を取得する
func FindLevels(db *gorm.DB) ([]string, error) {
	var levels []string
	if err := db.Table("logs").Distinct("level_name").Pluck("level_name", &levels).Error; err != nil {
		return nil, err
	}
	return levels, nil
}