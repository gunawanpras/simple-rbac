package handler

import (
	"errors"
)

func New(attr InitAttribute) Handler {
	if err := attr.validate(); err != nil {
		panic(err)
	}

	return &UserHandler{
		service: attr.Service,
	}
}

func (attr InitAttribute) validate() error {
	if attr.Service.UserService == nil {
		return errors.New("missing user service")
	}

	return nil
}
