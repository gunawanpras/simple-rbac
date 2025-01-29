package constant

import "net/http"

const (
	SUCCESS = "success"
	ERROR   = "error"
)

const (
	SYSTEM = "SYSTEM"
)

const (
	BindingParameterFailed = "failed to bind parameter"
)

const (
	UserCreateSuccess           = "user created successfully"
	UserCreateFailed            = "failed to create user"
	UserGetSuccess              = "user fetched successfully"
	UserGetFailed               = "failed to fetch user"
	UserNotFound                = "user not found"
	UserAlreadyExists           = "user already exists"
	RoleCreateSuccess           = "role created successfully"
	RoleCreateFailed            = "failed to create role"
	RoleGetSuccess              = "role fetched successfully"
	RoleGetFailed               = "failed to fetch role"
	RoleNotFound                = "role not found"
	RoleAlreadyExists           = "role already exists"
	PermissionCreateSuccess     = "permission created successfully"
	PermissionCreateFailed      = "failed to create permission"
	PermissionGetSuccess        = "permission fetched successfully"
	PermissionGetFailed         = "failed to fetch permission"
	PermissionNotFound          = "permission not found"
	PermissionAlreadyExists     = "permission already exists"
	RolePermissionCreateSuccess = "role permission created successfully"
	RolePermissionCreateFailed  = "failed to create role permission"
	RolePermissionGetSuccess    = "role permission fetched successfully"
	RolePermissionGetFailed     = "failed to fetch role permission"
	RolePermissionNotFound      = "role permission not found"
	RolePermissionAlreadyExists = "role permission already exists"
)

const (
	DbBeginTransactionFailed    = "failed to begin transaction: %v"
	DbRollbackTransactionFailed = "failed to rollback transaction: %v"
	DbCommitTransactionFailed   = "failed to commit transaction: %v"
	DbRecordNotFound            = "record not found"
)

var (
	GenericHttpStatusMappings = map[string]int{
		BindingParameterFailed: http.StatusBadRequest,
	}

	RbacHttpStatusMappings = map[string]int{
		UserCreateSuccess:           http.StatusCreated,
		UserCreateFailed:            http.StatusInternalServerError,
		UserGetSuccess:              http.StatusOK,
		UserGetFailed:               http.StatusInternalServerError,
		UserNotFound:                http.StatusNotFound,
		UserAlreadyExists:           http.StatusConflict,
		RoleCreateSuccess:           http.StatusCreated,
		RoleCreateFailed:            http.StatusInternalServerError,
		RoleGetSuccess:              http.StatusOK,
		RoleGetFailed:               http.StatusInternalServerError,
		RoleNotFound:                http.StatusNotFound,
		PermissionCreateSuccess:     http.StatusCreated,
		PermissionCreateFailed:      http.StatusInternalServerError,
		PermissionGetSuccess:        http.StatusOK,
		PermissionGetFailed:         http.StatusInternalServerError,
		PermissionNotFound:          http.StatusNotFound,
		RolePermissionCreateSuccess: http.StatusCreated,
		RolePermissionCreateFailed:  http.StatusInternalServerError,
		RolePermissionGetSuccess:    http.StatusOK,
		RolePermissionGetFailed:     http.StatusInternalServerError,
		RolePermissionNotFound:      http.StatusNotFound,
		RolePermissionAlreadyExists: http.StatusConflict,
		DbRecordNotFound:            http.StatusNotFound,
		DbBeginTransactionFailed:    http.StatusInternalServerError,
		DbRollbackTransactionFailed: http.StatusInternalServerError,
		DbCommitTransactionFailed:   http.StatusInternalServerError,
	}
)
