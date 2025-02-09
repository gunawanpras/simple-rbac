package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gunawanpras/simple-rbac/internal/setup"
)

func NewRouter(app *fiber.App, handler setup.Handler) {
	app.Use(recover.New())

	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return nil })
	app.Get("/docs", func(c *fiber.Ctx) error { return c.SendFile("./public/redoc.html") })

	users := app.Group("/users")
	users.Post("/", handler.UserHandler.CreateUser)
	users.Get("/", handler.UserHandler.ListUsers)
	users.Get("/:id", handler.UserHandler.GetUserByID)

	roles := app.Group("/roles")
	roles.Post("/", handler.RoleHandler.CreateRole)
	roles.Get("/", handler.RoleHandler.ListRoles)
	roles.Get("/:id", handler.RoleHandler.GetRoleByID)

	permissions := app.Group("/permissions")
	permissions.Post("/", handler.PermissionHandler.CreatePermission)
	permissions.Get("/", handler.PermissionHandler.ListPermissions)
	permissions.Get("/:id", handler.PermissionHandler.GetPermissionByID)

	rolePermissions := app.Group("/role-permissions")
	rolePermissions.Post("/", handler.RolePermissionHandler.CreateRolePermission)
	rolePermissions.Get("/", handler.RolePermissionHandler.ListRolePermissions)
	rolePermissions.Get("/:role_id", handler.RolePermissionHandler.ListRolePermissionsByRoleID)
	rolePermissions.Get("/:role_id/permissions/:permission_id", handler.RolePermissionHandler.GetRolePermissionByPermissionID)

	app.Static("/doc", "./public")
	app.Static("/docs/static", "./public/redoc.standalone.js")
}
