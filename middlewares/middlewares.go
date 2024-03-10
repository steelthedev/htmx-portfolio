package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func CheckHTMXInRequest() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Get("HX-Request") == "true" {
			c.Locals("IsHtmxRequest", true)
		} else {
			c.Locals("IsHtmxRequest", false)
		}
		return c.Next()
	}
}
