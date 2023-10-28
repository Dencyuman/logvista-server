package crud

import (
	"gorm.io/gorm"
)

func FindSystems(db *gorm.DB) ([]string, error) {
	var systems []string
	if err := db.Table("logs").Distinct("system_name").Pluck("system_name", &systems).Error; err != nil {
		return nil, err
	}
	return systems, nil
}
