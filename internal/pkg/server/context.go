package server

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

const (
	userKey = "user"
	connKey = "conn"
)

type (
	Context struct {
		echo.Context
	}
)

func (c *Context) GetAuthed() (*jwt.Token, bool) {
	val, ok := c.Get(userKey).(*jwt.Token)
	return val, ok
}

func setConn(c *Context, v interface{}) {
	c.Set(connKey, v)
}

func (c *Context) GetConn() (*gorm.DB, bool) {
	v, ok := c.Get(connKey).(*gorm.DB)
	return v, ok
}
