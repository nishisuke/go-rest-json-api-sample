package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUnauthedRoute(e *echo.Echo) {
	e.GET("/foo", func(c echo.Context) error {
		return c.String(http.StatusOK, "aaaaa")
	})
}
