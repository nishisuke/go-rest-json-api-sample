package api

import (
	"your_module/internal/app/api/infra/server"

	"github.com/labstack/echo/v4"
)

func Start(logger echo.Logger) error {
	if err := prepare(); err != nil {
		return err
	}

	return server.Start(logger)
}
