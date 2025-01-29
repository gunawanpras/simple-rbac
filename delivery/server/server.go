package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gunawanpras/simple-rbac/config"
	"github.com/gunawanpras/simple-rbac/helper/response"
	"github.com/gunawanpras/simple-rbac/util/constant"
)

func Up(handler Handler, config config.ServerConfig) {
	app := fiber.New(
		fiber.Config{
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				return response.Error(c, err.Error(), err, constant.GenericHttpStatusMappings)
			},
		},
	)
	app.Use(cors.New())

	NewRouter(app, handler)
	err := app.Listen(fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Panic(err)
	}
}
