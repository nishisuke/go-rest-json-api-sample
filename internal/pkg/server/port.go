package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func startOnPort(e *echo.Echo) error {
	v := os.Getenv("PORT")

	port, err := strconv.Atoi(v)
	if err != nil {
		return err
	}

	err = e.Start(fmt.Sprintf(":%d", port))
	return err
}
