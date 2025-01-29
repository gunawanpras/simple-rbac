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
	users.Post("/", handler.rbacHandler.CreateUser)
	users.Get("/", handler.rbacHandler.ListUsers)
	users.Get("/:id", handler.rbacHandler.GetUserByID)

	roles := app.Group("/roles")
	roles.Post("/", handler.rbacHandler.CreateRole)
	roles.Get("/", handler.rbacHandler.ListRoles)
	roles.Get("/:id", handler.rbacHandler.GetRoleByID)

	permissions := app.Group("/permissions")
	permissions.Post("/", handler.rbacHandler.CreatePermission)
	permissions.Get("/", handler.rbacHandler.ListPermissions)
	permissions.Get("/:id", handler.rbacHandler.GetPermissionByID)

	rolePermissions := app.Group("/role-permissions")
	rolePermissions.Post("/", handler.rbacHandler.CreateRolePermission)
	rolePermissions.Get("/", handler.rbacHandler.ListRolePermissions)
	rolePermissions.Get("/:role_id", handler.rbacHandler.ListRolePermissionsByRoleID)
	rolePermissions.Get("/:role_id/permissions/:permission_id", handler.rbacHandler.GetRolePermissionByPermissionID)

	app.Static("/doc", "./public")
	app.Static("/docs/static", "./public/redoc.standalone.js")
}
