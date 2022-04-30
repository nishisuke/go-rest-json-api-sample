package main

import (
	"your_module/internal/infra/logger"
	"your_module/internal/infra/server"
)

func main() {
	if err := server.Run(); err != nil {
		logger.NewLogger().FatalErr(err)
	}
}
