package main

import (
	"github.com/gofiber/template/html/v2"
	"github.com/steelthedev/htmx-portfolio/handler"
	"github.com/steelthedev/htmx-portfolio/middlewares"

	"github.com/gofiber/fiber/v2"
)

func main() {

	engine := html.New("./static", ".html")

	app := fiber.New(fiber.Config{
		Views:             engine,
		ViewsLayout:       "layouts/main",
		PassLocalsToViews: true,
	})

	app.Use(middlewares.CheckHTMXInRequest())

	app.Static("/static", "./assets")

	appHandler := handler.NewAppHander()

	app.Get("/", appHandler.Index)
	app.Get("/about", appHandler.About)
	app.Get("/contact", appHandler.Contact)
	app.Get("/blog", appHandler.Blog)
	app.Get("/portfolio", appHandler.Portfolio)
	app.Get("resume", appHandler.Resume)

	app.Listen(":3000")
}
