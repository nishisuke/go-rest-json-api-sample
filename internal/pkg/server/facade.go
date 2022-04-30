package server

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CallFacade[T any](c echo.Context, facade func(db *gorm.DB, ctx *Context, req T) error) error {
	req, err := ParseValidJSON[T](c)

	if err != nil {
		return err
	}

	ctx, ok := c.(*Context)
	if !ok {
		panic("Invalid type Context")
	}

	db, ok := ctx.GetConn()
	if !ok {
		panic("Invalid DB")
	}

	return facade(db, ctx, req)
}
