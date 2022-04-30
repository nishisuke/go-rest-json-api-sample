package server

import (
	"your_module/internal/app/api/infra/env"
	"your_module/internal/pkg/server"

	"github.com/labstack/echo/v4"
	// "github.com/nishisuke/echo-go1.18/internal/infra/env"
	// "github.com/nishisuke/echo-go1.18/internal/infra/httpreq"
	// "github.com/nishisuke/echo-go1.18/internal/infra/db"
)

func Start(logger echo.Logger, en env.Env) error {
	e := echo.New()
	e.Logger = logger

	e.HideBanner = !en.IsLocal()

	//e.HTTPErrorHandler = customHTTPErrorHandler
	//e.Validator = httpreq.NewValidation()
	// middlewares(e)
	// authed := e.Group("", auth)

	return server.StartOnPort(e)
}
