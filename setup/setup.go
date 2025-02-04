package setup

import (
	"github.com/gunawanpras/simple-rbac/config"
	"github.com/gunawanpras/simple-rbac/delivery/server"
	"github.com/jmoiron/sqlx"
)

type ExternalServices struct {
	Postgres *sqlx.DB
}

type CoreServices struct {
	Handler server.Handler
}

func InitExternalServices(conf *config.Config) *ExternalServices {
	pg := InitPostgres(conf)

	return &ExternalServices{
		Postgres: pg,
	}
}

func InitCoreServices(externalService *ExternalServices) *CoreServices {
	repository := server.NewRepository(externalService.Postgres)
	service := server.NewService(repository)
	handler := server.NewHandler(service)

	return &CoreServices{
		Handler: handler,
	}
}
