package main

import (
	"os"
	"your_module/internal/app/api"
	"your_module/internal/pkg/logger"
)

func main() {
	l := logger.Prepare(os.Stdout)
	if err := api.Start(l); err != nil {
		l.FatalErr(err)
	}
}
