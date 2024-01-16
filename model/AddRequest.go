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
type AddServiceRequest struct {
	ServiceName string `json:"service_name" validate:"required"`
	State       string `json:"status" validate:"required"`
	AddNodeRequest
	ServiceMetadata gormjsonb.JSONB `json:"service_metadata"`
}

type AddNodeRequest struct {
	NodeName     string          `json:"node_name" validate:"required"`
	NodeIp       string          `json:"node_ip" validate:"required"`
	NodeMetadata gormjsonb.JSONB `json:"node_metadata"`
}
