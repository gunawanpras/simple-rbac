package dto

import "github.com/google/uuid"

type CreatePermissionRequest struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"max=255"`
}

type GetPermissionByIDRequest struct {
	ID uuid.UUID `uri:"id" validate:"required,uuid"`
}
