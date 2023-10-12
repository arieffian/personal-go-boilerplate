package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type healthcheckHandler struct {
}

var _ HealthcheckService = (*healthcheckHandler)(nil)

func (h healthcheckHandler) HealthCheckHandler(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func NewHealthCheckHandler() *healthcheckHandler {
	return &healthcheckHandler{}
}
