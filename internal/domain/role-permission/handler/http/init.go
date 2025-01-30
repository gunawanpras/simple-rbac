package httphandler

import (
	"errors"

	"github.com/gunawanpras/simple-rbac/internal/domain/role-permission/handler"
)

func New(attr InitAttribute) handler.Handler {
	if err := attr.validate(); err != nil {
		panic(err)
	}

	return &RolePermissionHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if attr.Service.RolePermissionService == nil {
		return errors.New("missing role-permission service")
	}

	return nil
}
