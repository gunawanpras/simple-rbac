package main

import (
	"github.com/gunawanpras/simple-rbac/config"
	"github.com/gunawanpras/simple-rbac/delivery/server"
	"github.com/gunawanpras/simple-rbac/infrastructure/database/postgres"
)

func main() {
	// init config
	config.Init(config.WithConfigFile("config"), config.WithConfigType("yaml"))
	conf := config.Get()

	// init database
	postgre := postgres.NewInstance(conf)

	// init repo, service and handler
	repository := server.NewRepository(postgre)
	service := server.NewService(repository)
	handler := server.NewHandler(service)

	// init server
	server.Up(handler, conf.Server)
}
