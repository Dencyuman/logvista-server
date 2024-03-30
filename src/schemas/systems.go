package schemas

import (
	"time"
)

// swagger:request system
type System struct {
	Name     string `json:"name" example:"sample_system"`
	Category string `json:"category" example:"API Server"`
}

type SystemRequest struct {
	Category string `json:"category" example:"API Server"`
}

// swagger:response logResponse
type SystemResponse struct {
	ID string `json:"id" binding:"required" example:"00000000-0000-0000-0000-000000000000"`
	System
	CreatedAt time.Time `json:"created_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
	UpdatedAt time.Time `json:"updated_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
}

// swagger:response summary
type Summary struct {
	SystemResponse
	Data []SummaryData `json:"data" binding:"required"`
}

// swagger:model summaryData
type SummaryData struct {
	BaseTime        time.Time `json:"base_time" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
	InfologCount    int64     `json:"infolog_count" binding:"required" example:"10"`
	WarninglogCount int64     `json:"warninglog_count" binding:"required" example:"10"`
	ErrorlogCount   int64     `json:"errorlog_count" binding:"required" example:"10"`
}
