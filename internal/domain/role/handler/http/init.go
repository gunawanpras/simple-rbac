package httphandler

import (
	"errors"

	"github.com/gunawanpras/simple-rbac/internal/domain/role/handler"
)

func New(attr InitAttribute) handler.Handler {
	if err := attr.validate(); err != nil {
		panic(err)
	}

	return &RoleHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if attr.Service.RoleService == nil {
		return errors.New("missing role service")
	}

	return nil
}
