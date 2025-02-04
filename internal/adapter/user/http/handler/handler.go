package handler

import "github.com/gofiber/fiber/v2"

type Handler interface {
	CreateUser(c *fiber.Ctx) error
	ListUsers(c *fiber.Ctx) error
	GetUserByID(c *fiber.Ctx) error
}
