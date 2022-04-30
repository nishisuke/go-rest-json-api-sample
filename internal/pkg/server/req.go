package server

import "github.com/labstack/echo/v4"

func ParseValidJSON[T any](c echo.Context) (t T, err error) {
	if err := (&echo.DefaultBinder{}).BindBody(c, &t); err != nil {
		return t, err
	}

	if err := c.Validate(t); err != nil {
		return t, err
	}

	return t, nil
}
