package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Log struct {
	ID              string      `gorm:"type:uuid;primaryKey" json:"id"`
	SystemName      string      `json:"system_name"`
	CPUPercent      float64     `json:"cpu_percent"`
	ExcType         string      `json:"exc_type"`
	ExcValue        string      `json:"exc_value"`
	ExcDetail       string      `json:"exc_detail"`
	ExcTraceback    []Traceback `gorm:"foreignKey:LogID" json:"exc_traceback"`
	FileName        string      `json:"file_name"`
	FuncName        string      `json:"func_name"`
	Lineno          int         `json:"lineno"`
	Message         string      `json:"message"`
	Module          string      `json:"module"`
	Name            string      `json:"name"`
	LevelName       string      `json:"level_name"`
	Levelno         int         `json:"levelno"`
	Process         int         `json:"process"`
	ProcessName     string      `json:"process_name"`
	Thread          int         `json:"thread"`
	ThreadName      string      `json:"thread_name"`
	TotalMemory     int64       `json:"total_memory"`
	AvailableMemory int64       `json:"available_memory"`
	MemoryPercent   float64     `json:"memory_percent"`
	UsedMemory      int64       `json:"used_memory"`
	FreeMemory      int64       `json:"free_memory"`
	CPUUserTime     float64     `json:"cpu_user_time"`
	CPUSystemTime   float64     `json:"cpu_system_time"`
	CPUIdleTime     float64     `json:"cpu_idle_time"`
	Timestamp       time.Time   `json:"timestamp"`
	Attributes      string      `gorm:"type:text" json:"attributes"`
	CreatedAt       time.Time   `gorm:"autoCreateTime"`
	UpdatedAt       time.Time   `gorm:"autoUpdateTime"`
}

type Traceback struct {
	ID         string    `gorm:"type:uuid;primaryKey" json:"id"`
	LogID      string    `json:"log_id"`
	TbFilename string    `json:"tb_filename"`
	TbLineno   int       `json:"tb_lineno"`
	TbName     string    `json:"tb_name"`
	TbLine     string    `json:"tb_line"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

func (t *Traceback) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New().String()
	return
}
