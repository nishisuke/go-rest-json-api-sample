package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nishisuke/ezlog"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Start(logger echo.Logger, v interface{}, cb func(e *echo.Echo)) error {
	e := echo.New()
	ezlog.Prepare(e)

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{c}
			setConn(cc, v)
			return next(cc)
		}
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	cb(e)

	return startOnPort(e)
}
