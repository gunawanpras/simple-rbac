package main

import (
	"github.com/gunawanpras/simple-rbac/config"
	"github.com/gunawanpras/simple-rbac/delivery/server"
	"github.com/gunawanpras/simple-rbac/internal/setup"
)

func main() {
	// init config
	config.Init(config.WithConfigFile("config"), config.WithConfigType("yaml"))
	conf := config.Get()

	// init external services
	externalService := setup.InitExternalServices(conf)
	defer externalService.Postgres.Close()

	// init core services
	coreService := setup.InitCoreServices(externalService)

	// init server
	server.Up(coreService.Handler, conf.Server)
}
