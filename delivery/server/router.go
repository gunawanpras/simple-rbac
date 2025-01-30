package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func NewRouter(app *fiber.App, handler Handler) {
	app.Use(recover.New())

	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return nil })
	app.Get("/docs", func(c *fiber.Ctx) error { return c.SendFile("./public/redoc.html") })

	users := app.Group("/users")
	users.Post("/", handler.userHandler.CreateUser)
	users.Get("/", handler.userHandler.ListUsers)
	users.Get("/:id", handler.userHandler.GetUserByID)

	roles := app.Group("/roles")
	roles.Post("/", handler.roleHandler.CreateRole)
	roles.Get("/", handler.roleHandler.ListRoles)
	roles.Get("/:id", handler.roleHandler.GetRoleByID)

	permissions := app.Group("/permissions")
	permissions.Post("/", handler.permissionHandler.CreatePermission)
	permissions.Get("/", handler.permissionHandler.ListPermissions)
	permissions.Get("/:id", handler.permissionHandler.GetPermissionByID)

	rolePermissions := app.Group("/role-permissions")
	rolePermissions.Post("/", handler.rolePermissionHandler.CreateRolePermission)
	rolePermissions.Get("/", handler.rolePermissionHandler.ListRolePermissions)
	rolePermissions.Get("/:role_id", handler.rolePermissionHandler.ListRolePermissionsByRoleID)
	rolePermissions.Get("/:role_id/permissions/:permission_id", handler.rolePermissionHandler.GetRolePermissionByPermissionID)

	app.Static("/doc", "./public")
	app.Static("/docs/static", "./public/redoc.standalone.js")
}
