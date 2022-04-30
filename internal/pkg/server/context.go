package server

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

const key = "user"

type (
	Context struct {
		echo.Context
	}
)

func (c *Context) GetAuthed() (*jwt.Token, bool) {
	val, ok := c.Get(key).(*jwt.Token)
	return val, ok
}
