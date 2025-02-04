package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreatePermission(c *fiber.Ctx) error
	ListPermissions(c *fiber.Ctx) error
	GetPermissionByID(c *fiber.Ctx) error
}
