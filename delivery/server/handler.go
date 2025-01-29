package server

import (
	"github.com/gunawanpras/simple-rbac/internal/domain/rbac/handler"
	rbac "github.com/gunawanpras/simple-rbac/internal/domain/rbac/handler/http"
)

type Handler struct {
	rbacHandler handler.Handler
}

func NewHandler(service Service) Handler {
	return Handler{
		rbacHandler: rbac.New(rbac.InitAttribute{
			Service: rbac.ServiceAttribute{
				RbacService: service.RbacService,
			},
		}),
	}
}
