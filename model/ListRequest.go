package model

type ListRequest struct {
	Id int64 `json:"id" validate:"required"`
	AddRequest
}
type ListServiceRequest struct {
	ServiceId   int64           `json:"id" validate:"required"`
	ServiceName string          `json:"service_name" validate:"required"`
	State       string          `json:"status" validate:"required"`
	UpTime      string          `json:"uptime"`
	StartTime   string          `json:"start_time" validate:"required"`
	EndTime     string          `json:"end_time"`
	Node        ListNodeRequest `json:"node"`
}

type ListNodeRequest struct {
	NodeId int64 `json:"id" validate:"required"`
	AddNodeRequest
}
