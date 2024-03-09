package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/steelthedev/htmx-portfolio/view/app"
)

type AppHandler struct{}

func (h AppHandler) HandleIndex(c echo.Context) error {
	return render(c, app.Index())
}

func (h AppHandler) HandleAbout(c echo.Context) error {
	return render(c, app.About())
}
