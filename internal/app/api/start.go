package api

import (
	"your_module/internal/app/api/err"
	"your_module/internal/app/api/infra/env"
	"your_module/internal/app/api/infra/server"
	"your_module/internal/app/api/routes"
	"your_module/internal/app/api/validation"
	"your_module/internal/pkg/db"

	"github.com/labstack/echo/v4"
)

func Start(logger echo.Logger, en env.Env) error {
	conn, e := db.Prepare(en.GormLogger())
	if e != nil {
		return e
	}

	e = server.Start(logger, conn, validation.NewValidation(), err.Resopond, en.IsLocal(), routes.RegisterUnauthedRoute)
	return e
}
