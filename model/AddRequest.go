package model

import gormjsonb "github.com/dariubs/gorm-jsonb"

type AddRequest struct {
	ServiceName string          `json:"service_name" validate:"required"`
	State       string          `json:"status" validate:"required"`
	UpTime      string          `json:"uptime"`
	StartTime   string          `json:"start_time" validate:"required"`
	EndTime     string          `json:"end_time"`
	Metadata    gormjsonb.JSONB `json:"metadata"`
}
