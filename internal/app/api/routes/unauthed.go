package routes

import (
	"net/http"
	"time"
	"your_module/internal/pkg/server"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterUnauthedRoute(e *echo.Group) {
	e.POST("/foo", func(c echo.Context) error {
		return server.CallFacade(c, (&Controller{}).Hoge)
	})
}

type Controller struct{}

// Hoge godoc
// @Summary      List accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   User
// @Failure      400  {object}  Error
// @Failure      404  {object}  Error
// @Failure      500  {object}  Error
// @Router       /accounts [get]
func (cont *Controller) Hoge(db *gorm.DB, ctx *server.Context, req User) error {
	return ctx.String(http.StatusOK, "aaaaa")
}

type (
	User struct {
		ID        uint
		Name      string `json:"name" validate:"required"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
	Error struct {
		Name string `json:"name" validate:"required"`
	}
)
