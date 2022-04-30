package main

import (
	"your_module/internal/app/api"
	"your_module/internal/pkg/logger"
)

func main() {
	l := logger.Prepare()
	if err := api.Start(l); err != nil {
		l.FatalErr(err)
	}
}
