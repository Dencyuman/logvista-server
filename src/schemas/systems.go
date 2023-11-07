package schemas

import (
	"time"
)

// swagger:request system
type System struct {
	Name       string    `json:"name" example:"sample_system"`
	Category   string    `json:"category" example:"API Server"`
}

// swagger:response logResponse
type SystemResponse struct {
	ID              string           `json:"id" binding:"required" example:"00000000-0000-0000-0000-000000000000"`
	System
	CreatedAt    time.Time           `json:"created_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
	UpdatedAt    time.Time           `json:"updated_at" binding:"required" example:"2023-01-01T00:00:00.000000+09:00"`
}