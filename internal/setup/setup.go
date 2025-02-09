package setup

import (
	"github.com/gunawanpras/simple-rbac/config"
	"github.com/jmoiron/sqlx"
)

type ExternalServices struct {
	Postgres *sqlx.DB
}

type CoreServices struct {
	Handler Handler
}

func InitExternalServices(conf *config.Config) *ExternalServices {
	pg := InitPostgres(conf)

	return &ExternalServices{
		Postgres: pg,
	}
}

func InitCoreServices(externalService *ExternalServices) *CoreServices {
	repository := NewRepository(externalService.Postgres)
	service := NewService(repository)
	handler := NewHandler(service)

	return &CoreServices{
		Handler: handler,
	}
}
