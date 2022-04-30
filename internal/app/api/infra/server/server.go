package server

// CHECK

import (
	"your_module/internal/app/api/err"
	"your_module/internal/pkg/server"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const expectedContentType = "application/json"

func Start(logger echo.Logger, validator echo.Validator, handler echo.HTTPErrorHandler, isLocal bool, register func(e *echo.Echo)) error {
	err := server.Start(logger, func(e *echo.Echo) {
		e.HideBanner = !isLocal

		e.Validator = validator

		e.HTTPErrorHandler = handler

		middlewares(e)

		register(e)
	})
	return err
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
	// e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
	//	XSSProtection:         "",
	//	ContentTypeNosniff:    "",
	//	XFrameOptions:         "",
	//	HSTSMaxAge:            3600,
	//	ContentSecurityPolicy: "default-src 'self'",
	//
	// }))
}
