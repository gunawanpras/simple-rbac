package dto

import "github.com/google/uuid"

type CreateRoleRequest struct {
	Name string `json:"name" validate:"required,min=3,max=100"`
}

type GetRoleByIDRequest struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}
