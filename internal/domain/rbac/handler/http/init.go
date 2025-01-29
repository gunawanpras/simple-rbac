package httphandler

import (
	"errors"

	"github.com/gunawanpras/simple-rbac/internal/domain/rbac/handler"
)

func New(attr InitAttribute) handler.Handler {
	if err := attr.validate(); err != nil {
		panic(err)
	}

	return &RbacHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if attr.Service.RbacService == nil {
		return errors.New("missing estate service")
	}

	return nil
}
