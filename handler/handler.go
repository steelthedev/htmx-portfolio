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

func (h AppHandler) Blog(c *fiber.Ctx) error {
	return c.Render("blog", fiber.Map{})
}

func (h AppHandler) Contact(c *fiber.Ctx) error {
	return c.Render("contact", fiber.Map{})
}

func (h AppHandler) Portfolio(c *fiber.Ctx) error {
	return c.Render("portfolio", fiber.Map{})
}

func (h AppHandler) Resume(c *fiber.Ctx) error {
	return c.Render("resume", fiber.Map{})
}
