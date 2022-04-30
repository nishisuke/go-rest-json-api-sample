package server

// CHECK

import (
	"your_module/internal/pkg/server"

	"github.com/labstack/echo/v4"
	// "github.com/nishisuke/echo-go1.18/internal/infra/env"
	// "github.com/nishisuke/echo-go1.18/internal/infra/httpreq"
	// "github.com/nishisuke/echo-go1.18/internal/infra/db"
)

func Start(logger echo.Logger) error {
	e := echo.New()
	e.Logger = logger

	//e.Logger.SetLevel(envi.LogLevel())
	//e.HideBanner = !envi.IsLocal()

	//e.HTTPErrorHandler = customHTTPErrorHandler
	//e.Validator = httpreq.NewValidation()
	// middlewares(e)
	// authed := e.Group("", auth)

	return server.StartOnPort(e)
}
