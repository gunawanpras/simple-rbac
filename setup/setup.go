package setup

import (
	"github.com/gunawanpras/simple-rbac/config"
	"github.com/jmoiron/sqlx"
)

type ExternalServices struct {
	Postgres *sqlx.DB
}

func InitExternalServices(conf *config.Config) *ExternalServices {
	pg := InitPostgres(conf)

	return &ExternalServices{
		Postgres: pg,
	}
}
