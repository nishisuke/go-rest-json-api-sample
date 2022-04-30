package api

import (
	"your_module/internal/app/api/infra/server"
	"your_module/internal/pkg/db"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var gormDB *gorm.DB

func Start(logger echo.Logger) error {
	env := prepare()

	tmp, err := db.Prepare(env.GormLogger())
	if err != nil {
		return err
	}
	gormDB = tmp

	return server.Start(logger)
}
