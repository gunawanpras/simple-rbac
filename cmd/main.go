package main

import (
	"github.com/gunawanpras/simple-rbac/config"
	"github.com/gunawanpras/simple-rbac/delivery/server"
	"github.com/gunawanpras/simple-rbac/setup"
)

func main() {
	// init config
	config.Init(config.WithConfigFile("config"), config.WithConfigType("yaml"))
	conf := config.Get()

	// init external services
	externalService := setup.InitExternalServices(conf)
	defer externalService.Postgres.Close()

	// init repo, service and handler
	repository := server.NewRepository(externalService.Postgres)
	service := server.NewService(repository)
	handler := server.NewHandler(service)

	// init server
	server.Up(handler, conf.Server)
}
