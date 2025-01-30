package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateRole(c *fiber.Ctx) error
	ListRoles(c *fiber.Ctx) error
	GetRoleByID(c *fiber.Ctx) error
}
