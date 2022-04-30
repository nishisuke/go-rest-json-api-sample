package main

import (
	"your_module/internal/pkg/logger"
	"your_module/internal/pkg/server"
)

func main() {
	if err := server.Run(); err != nil {
		logger.DefaultLogger().FatalErr(err)
	}
}
