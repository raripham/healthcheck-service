package model

import gormjsonb "github.com/dariubs/gorm-jsonb"

type ListRequest struct {
	Id int64 `json:"id" validate:"required"`
	AddRequest
	NodeId int64 `json:"node_id" validate:"required"`
}
type ListServiceRequest struct {
	ServiceId       int64           `json:"service_id" validate:"required"`
	ServiceName     string          `json:"service_name" validate:"required"`
	State           string          `json:"status" validate:"required"`
	UpTime          string          `json:"uptime"`
	StartTime       string          `json:"start_time" validate:"required"`
	EndTime         string          `json:"end_time"`
	Node            ListNodeRequest `json:"node"`
	ServiceMetadata gormjsonb.JSONB `json:"service_metadata"`
	// Node            ListNodeRequest `json:"node"`
}

type ListNodeRequest struct {
	NodeId int64 `json:"node_id" validate:"required"`
	AddNodeRequest
}
