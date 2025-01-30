package model

import "github.com/google/uuid"

type CreatePermissionResponse struct {
	ID uuid.UUID `json:"id"`
}

type GetPermissionResponse struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedBy string    `json:"created_by"`
	CreatedAt string    `json:"created_at"`
}

type ListPermissionsResponse []GetPermissionResponse

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ServiceResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
