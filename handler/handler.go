package handler

import "github.com/gofiber/fiber/v2"

type AppHandler struct {
}

func NewAppHander() *AppHandler {
	return &AppHandler{}
}

func (h AppHandler) Index(c *fiber.Ctx) error {

	return c.Render("index", fiber.Map{})
}

func (h AppHandler) About(c *fiber.Ctx) error {
	return c.Render("about", fiber.Map{})
}
