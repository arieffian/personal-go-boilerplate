package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// New creates a new middleware handler
func NewValidateAPIKey(apiKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		rApiKey := c.Request().Header.Peek("X-API-KEY")

		if rApiKey != nil && string(rApiKey) == apiKey {
			return c.Next()
		}

		// Authentication failed
		return c.SendStatus(fiber.StatusUnauthorized)
	}
}
