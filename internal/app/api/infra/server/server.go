package server

import (
	"your_module/internal/app/api/err"
	"your_module/internal/app/api/infra/env"
	"your_module/internal/pkg/server"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	// "github.com/nishisuke/echo-go1.18/internal/infra/env"
	// "github.com/nishisuke/echo-go1.18/internal/infra/httpreq"
	// "github.com/nishisuke/echo-go1.18/internal/infra/db"
)

const expectedContentType = "application/json"

type (
	Context struct {
		echo.Context
	}
)

func (c *Context) GetAuthed() *jwt.Token {
	return c.Get("user").(*jwt.Token)
}
func Start(logger echo.Logger, validator echo.Validator, handler echo.HTTPErrorHandler, en env.Env, register func(e *echo.Echo)) error {
	e := echo.New()
	e.Logger = logger
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &Context{c}
			return next(cc)
		}
	})

	e.HideBanner = !en.IsLocal()

	e.Validator = validator

	e.HTTPErrorHandler = handler

	middlewares(e)
	// authed := e.Group("", auth)

	register(e)
	return server.StartOnPort(e)
}

func middlewares(e *echo.Echo) {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"}, // TODO
	}))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if c.Request().Header.Get("content-type") != expectedContentType {
				return err.ErrInvalidContentType
			}

			q := c.QueryString()
			if q != "" {
				return err.ErrQueryParamsExist
			}

			return next(c)

		}
	})

	// TODO
	// e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
	// 	SigningKey: []byte("secret"),
	// 	KeyFunc:    getKey,

	// 	// TokenLookup: "query:token",
	// 	ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
	// 		keyFunc := func(t *jwt.Token) (interface{}, error) {
	// 			if t.Method.Alg() != "HS256" {
	// 				return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	// 			}
	// 			return []byte("secret"), nil
	// 		}

	// 		// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
	// 		token, err := jwt.Parse(auth, keyFunc)
	// 		return &jwt.Token{}, nil // TODO: Rm this line. this is danger.
	// 		if err != nil {
	// 			fmt.Println("aaaaaaa2", err)
	// 			return nil, err
	// 		}

	// 		if !token.Valid {
	// 			return nil, errors.New("invalid token")
	// 		}
	// 		return token, nil
	// 	},
	// }))

	// TODO
	//e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	//	XSSProtection:         "",
	//	ContentTypeNosniff:    "",
	//	XFrameOptions:         "",
	//	HSTSMaxAge:            3600,
	//	ContentSecurityPolicy: "default-src 'self'",
	//
	//}))
}
