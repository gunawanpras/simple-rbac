package postgres

import (
	"log"
	"time"

	"github.com/gunawanpras/simple-rbac/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewInstance(conf *config.Config) *sqlx.DB {
	pg, err := sqlx.Connect("postgres", conf.Postgre.Rbac.ReadWrite.ConnString)
	if err != nil {
		log.Panic("failed to open postgre client for rbac service:", err)
	}

	pg.SetMaxOpenConns(conf.Postgre.Rbac.ReadWrite.MaxOpenConn)
	pg.SetMaxIdleConns(conf.Postgre.Rbac.ReadWrite.MaxIdleConn)
	pg.SetConnMaxLifetime(time.Second * time.Duration(conf.Postgre.Rbac.ReadWrite.MaxConnLifeTimeInSecond))

	return pg
}
