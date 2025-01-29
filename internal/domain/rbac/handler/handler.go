package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateUser(c *fiber.Ctx) error
	ListUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
	CreateRole(c *fiber.Ctx) error
	ListRoles(c *fiber.Ctx) error
	GetRoleByID(c *fiber.Ctx) error
	CreatePermission(c *fiber.Ctx) error
	ListPermissions(c *fiber.Ctx) error
	GetPermissionByID(c *fiber.Ctx) error
	CreateRolePermission(c *fiber.Ctx) error
	ListRolePermissions(c *fiber.Ctx) error
	ListRolePermissionsByRoleID(c *fiber.Ctx) error
	GetRolePermissionByPermissionID(c *fiber.Ctx) error
}
