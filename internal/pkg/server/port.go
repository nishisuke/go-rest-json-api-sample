package server

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
)

func getPort() (int, error) {
	v := os.Getenv("PORT")
	return strconv.Atoi(v)
}

func StartOnPort(e *echo.Echo) error {
	port, err := getPort()
	if err != nil {
		return err
	}

	return e.Start(fmt.Sprintf(":%d", port))
}
