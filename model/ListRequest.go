package model

type ListRequest struct {
	Id int64 `json:"id" validate:"required"`
	AddRequest
}
