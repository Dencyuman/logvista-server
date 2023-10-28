package crud

import (
	"gorm.io/gorm"
)

// 存在するログレベル一覧を取得する
func FindLevels(db *gorm.DB) ([]string, error) {
	var levels []string
	if err := db.Table("logs").Distinct("level_name").Pluck("level_name", &levels).Error; err != nil {
		return nil, err
	}
	return levels, nil
}

// 存在するシステム名一覧を取得する
func FindSystems(db *gorm.DB) ([]string, error) {
	var systems []string
	if err := db.Table("logs").Distinct("system_name").Pluck("system_name", &systems).Error; err != nil {
		return nil, err
	}
	return systems, nil
}
