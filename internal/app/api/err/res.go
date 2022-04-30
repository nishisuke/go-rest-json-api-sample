package err

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Resopond(e error, c echo.Context) {
	code, msg := codeAndMessage(e)
	if msg == "" {
		msg = http.StatusText(code)
	}

	err := c.JSON(code, map[string]interface{}{
		"message": msg,
	})

	if err != nil {
		c.Logger().Error(err)
	}
}

func codeAndMessage(e error) (int, string) {
	code := http.StatusInternalServerError

	var echoErr *echo.HTTPError
	if errors.As(e, &echoErr) {
		return echoErr.Code, fmt.Sprint(echoErr.Message)
	}

	var apiErr APIError
	if errors.As(e, &apiErr) {
		return apiErr.HTTPStatus(), apiErr.Error()
	}

	return code, ""
}
