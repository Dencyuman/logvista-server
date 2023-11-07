package schemas

import (
	"time"
)

// swagger:request flexibleLog
type FlexibleLog struct {
	ID         string                 `json:"id" binding:"required" example:"00000000-0000-0000-0000-000000000000"`
	Timestamp  time.Time              `json:"timestamp" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
	Level      string                 `json:"level" binding:"required" example:"INFO"`
	Attributes map[string]interface{} `json:"attributes"`
}

// swagger:request log
type Log struct {
	ID              string                 `json:"id" binding:"required" example:"00000000-0000-0000-0000-000000000000"`
	SystemName      string                 `json:"system_name" example:"sample_system"`
	CPUPercent      float64                `json:"cpu_percent" example:"0.0"`
	ExcType         string                 `json:"exc_type" example:"Exception"`
	ExcValue        string                 `json:"exc_value" example:"sample"`
	ExcDetail       string                 `json:"exc_detail" example:"Some traceback details here"`
	ExcTraceback    []Traceback            `json:"exc_traceback"`
	FileName        string                 `json:"file_name" example:"sample.py"`
	FuncName        string                 `json:"func_name" example:"<module>"`
	Lineno          int                    `json:"lineno" example:"1"`
	Message         string                 `json:"message" example:"sample message"`
	Module          string                 `json:"module" example:"sample"`
	Name            string                 `json:"name" example:"sample"`
	LevelName       string                 `json:"level_name" binding:"required" example:"INFO"`
	Levelno         int                    `json:"levelno" example:"20"`
	Process         int                    `json:"process" example:"1000"`
	ProcessName     string                 `json:"process_name" example:"MainProcess"`
	Thread          int                    `json:"thread" example:"10000"`
	ThreadName      string                 `json:"thread_name" example:"MainThread"`
	TotalMemory     int64                  `json:"total_memory" example:"16000000"`
	AvailableMemory int64                  `json:"available_memory" example:"8000000"`
	MemoryPercent   float64                `json:"memory_percent" example:"50.0"`
	UsedMemory      int64                  `json:"used_memory" example:"8000000"`
	FreeMemory      int64                  `json:"free_memory" example:"8000000"`
	CPUUserTime     float64                `json:"cpu_user_time" example:"700.0"`
	CPUSystemTime   float64                `json:"cpu_system_time" example:"500.0"`
	CPUIdleTime     float64                `json:"cpu_idle_time" example:"10000.0"`
	Timestamp       time.Time              `json:"timestamp" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
	Attributes      map[string]interface{} `json:"attributes"`
}

// swagger:model traceback
type Traceback struct {
	TbFilename string `json:"tb_filename" example:"C:\\User\\USER\\sample.py"`
	TbLineno   int    `json:"tb_lineno" example:"31"`
	TbName     string `json:"tb_name" example:"<module>"`
	TbLine     string `json:"tb_line" example:"raise Exception(\"sample\")"`
}

// swagger:response logResponse
type LogResponse struct {
	Log
	CreatedAt    time.Time           `json:"created_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
	UpdatedAt    time.Time           `json:"updated_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
	ExcTraceback []TracebackResponse `json:"exc_traceback"`
}

// swagger:model tracebackResponse
type TracebackResponse struct {
	Traceback
}

// swagger:response paginatedLogResponse
type PaginatedLogResponse struct {
	Total      int           `json:"total" example:"100"`      // 総アイテム数
	Page       int           `json:"page" example:"1"`         // 現在のページ数
	Limit      int           `json:"limit" example:"10"`       // 1ページあたりのアイテム数
	TotalPages int           `json:"total_pages" example:"10"` // 総ページ数
	Items      []LogResponse `json:"items"`                    // ログの配列
}

// swagger:response summary
type Summary struct {
	ID string `json:"id" binding:"required" example:"00000000-0000-0000-0000-000000000000"`
	System
	LatestLog LogResponse   `json:"latest_log" binding:"required"`
	Data      []SummaryData `json:"data" binding:"required"`
}

// swagger:model summaryData
type SummaryData struct {
	BaseTime        time.Time `json:"base_time" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
	InfologCount    int64     `json:"infolog_count" binding:"required" example:"10"`
	WarninglogCount int64     `json:"warninglog_count" binding:"required" example:"10"`
	ErrorlogCount   int64     `json:"errorlog_count" binding:"required" example:"10"`
}
