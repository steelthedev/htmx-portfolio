package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/steelthedev/htmx-portfolio/handler"
)

func main() {
	app := echo.New()
	app.Static("/", "assets")

	appHandler := handler.AppHandler{}

	app.GET("/", appHandler.HandleIndex)
	app.GET("/about", appHandler.HandleAbout)

	fmt.Println("starting server on port 3000")
	app.Start(":3000")
}
