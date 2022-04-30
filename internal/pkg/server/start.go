package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Start(logger echo.Logger, cb func(e *echo.Echo)) error {
	e := echo.New()
	e.Logger = logger

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{c}
			return next(cc)
		}
	})

	cb(e)

	return startOnPort(e)
}
