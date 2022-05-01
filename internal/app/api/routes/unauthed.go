package routes

import (
	"net/http"
	"time"
	"your_module/internal/pkg/server"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type (
	User struct {
		ID        uint
		Name      string `json:"name" validate:"required"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

func RegisterUnauthedRoute(e *echo.Echo) {
	e.POST("/foo", func(c echo.Context) error {
		return server.CallFacade(c, func(db *gorm.DB, ctx *server.Context, req User) error {
			var u User
			err := db.First(&u).Error
			if err != nil {
				return err
			}
			c.Logger().Debug(u, req)
			return c.String(http.StatusOK, "aaaaa")
		})
	})
}
