package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type HealthcheckConfigType string

const (
	SiteTitle HealthcheckConfigType = "SiteTitle"
	Endpoint  HealthcheckConfigType = "Endpoint"
)

type HealthcheckConfig struct {
	ID            string                `gorm:"type:uuid;primaryKey" json:"id"`
	SystemID      string                `gorm:"type:uuid;index" json:"system_id"`
	System        System                `gorm:"foreignKey:SystemID" json:"system"`
	Name          string                `json:"name"`
	Description   string                `json:"description"`
	ConfigType    HealthcheckConfigType `json:"config_type"`
	ExpectedValue string                `json:"expected_value"`
	Url           string                `json:"url"`
	Timespan      int                   `json:"timespan"`
	Note          string                `json:"note"`
	CreatedAt     time.Time             `gorm:"autoCreateTime"`
	UpdatedAt     time.Time             `gorm:"autoUpdateTime"`
}

func (h *HealthcheckConfig) BeforeCreate(tx *gorm.DB) (err error) {
	h.ID = uuid.New().String()
	return
}