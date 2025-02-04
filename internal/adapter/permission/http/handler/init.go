package handler

import (
	"errors"
)

func New(attr InitAttribute) Handler {
	if err := attr.validate(); err != nil {
		panic(err)
	}

	return &PermissionHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if attr.Service.PermissionService == nil {
		return errors.New("missing permission service")
	}

	return nil
}
