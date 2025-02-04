package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateRolePermission(c *fiber.Ctx) error
	ListRolePermissions(c *fiber.Ctx) error
	ListRolePermissionsByRoleID(c *fiber.Ctx) error
	GetRolePermissionByPermissionID(c *fiber.Ctx) error
}
