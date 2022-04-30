package api

import (
	"your_module/internal/app/api/infra/env"
	"your_module/internal/app/api/infra/server"
	"your_module/internal/app/api/validation"
	"your_module/internal/pkg/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func Start(logger echo.Logger, e env.Env) error {
	tmp, err := db.Prepare(e.GormLogger())
	if err != nil {
		return err
	}
	gormDB = tmp

	return server.Start(logger, validation.NewValidation(), e)
}
