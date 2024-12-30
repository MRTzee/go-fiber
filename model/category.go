package model

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}
