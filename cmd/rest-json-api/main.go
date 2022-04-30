package main

import (
	"os"
	"your_module/internal/app/api"
	"your_module/internal/pkg/env"
	"your_module/internal/pkg/logger"
)

func main() {
	e := env.Prepare()
	l := logger.Prepare(os.Stdout, e.LogLevel())

	if err := api.Start(l, e); err != nil {
		l.FatalErr(err)
	}
}
