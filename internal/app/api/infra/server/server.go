package server

import (
	"your_module/internal/pkg/server"

	"github.com/labstack/echo/v4"
	// "github.com/nishisuke/echo-go1.18/internal/infra/env"
	// "github.com/nishisuke/echo-go1.18/internal/infra/httpreq"
	// "github.com/nishisuke/echo-go1.18/internal/infra/db"
)

var infras = []func() error{
	// db.prepare,
}

func Start(logger echo.Logger) error {
	return prepare(func() error {
		e := echo.New()
		e.Logger = logger

		//envi := env.Env()

		//e.Logger.SetLevel(envi.LogLevel())
		//e.HideBanner = !envi.IsLocal()

		//e.HTTPErrorHandler = customHTTPErrorHandler
		//e.Validator = httpreq.NewValidation()
		// middlewares(e)
		// authed := e.Group("", auth)

		return server.StartOnPort(e)
	})
}
func prepare(cb func() error) error {
	for _, infra := range infras {
		if err := infra(); err != nil {
			return err
		}
	}

	return cb()
}

// func callFacade[T any](c echo.Context, cb func(c *Context, req T) error) error {
// 	a, err := httpreq.ParseValidJSON[T](c)
//
// 	if err != nil {
// 		return err
// 	}
//
// 	foo, ok := c.(*Context)
//
// 	if !ok {
// 		panic("context invalid")
// 	}
//
// 	return cb(foo, *a)
// }
//
// type (
// 	Context struct {
// 		echo.Context
// 	}
// )
//
// func (c *Context) GetAuthed() *jwt.Token {
// 	return c.Get("user").(*jwt.Token)
// }
//// TODO: Impl auth
//func auth(next echo.HandlerFunc) echo.HandlerFunc {
//	return func(c echo.Context) error {
//		c.Logger().Debugj(map[string]interface{}{
//			"msg": "AUTHED",
//		})
//		return next(c)
//	}
//}
//
