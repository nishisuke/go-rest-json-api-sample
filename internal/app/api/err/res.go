package err

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Resopond(e error, c echo.Context) {
	code, msg := codeAndMessage(e)

	c.JSON(code, map[string]interface{}{
		"message": msg,
	})
}

func codeAndMessage(e error) (int, string) {
	code := http.StatusInternalServerError
	msg := ""

	switch v := e.(type) {
	case *echo.HTTPError:
		code = v.Code
		msg = fmt.Sprint(v.Message)
	case APIErr:
		code = v.HttpStatus()
		msg = v.Error()
	}
	if msg == "" {
		msg = http.StatusText(code)
	}
	return code, msg
}
